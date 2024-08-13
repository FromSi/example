package database

import (
	"errors"
	"github.com/fromsi/example/cmd/apiserver/config"
	"github.com/fromsi/example/cmd/apiserver/types"
	domairepository "github.com/fromsi/example/internal/app/apiserver/domain/repositories"
	"github.com/fromsi/example/internal/app/apiserver/infrastructure/models"
	"github.com/fromsi/example/internal/app/apiserver/infrastructure/repositories"
	"go.uber.org/fx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MasterGormDB *gorm.DB
type SlaveGormDB *gorm.DB
type TestGormDB *gorm.DB

type RelationDatabase struct {
	MasterORMType        string
	MasterConnectionType string
	SlaveORMType         string
	SlaveConnectionType  string
	TestConnectionType   string
	Config               config.Config
}

func NewRelationDatabase(applicationConfig *config.Config) (*RelationDatabase, error) {
	masterORMType := applicationConfig.RelationDatabase.Connection.MasterORM
	correctMasterORMType := masterORMType == types.DatabaseGormORMType

	if !correctMasterORMType {
		return nil, errors.New("master orm type is incorrect")
	}

	masterConnectionType := applicationConfig.RelationDatabase.Connection.Master
	correctMasterConnectionType := masterConnectionType == types.DatabaseSQLiteDatabaseType

	if !correctMasterConnectionType {
		return nil, errors.New("master connection type is incorrect")
	}

	slaveORMType := applicationConfig.RelationDatabase.Connection.SlaveORM
	correctSlaveORMType := slaveORMType == types.DatabaseGormORMType

	if !correctSlaveORMType {
		return nil, errors.New("slave orm type is incorrect")
	}

	slaveConnectionType := applicationConfig.RelationDatabase.Connection.Slave
	correctSlaveConnectionType := slaveConnectionType == types.DatabaseSQLiteDatabaseType

	if !correctSlaveConnectionType {
		return nil, errors.New("slave connection type is incorrect")
	}

	testConnectionType := applicationConfig.RelationDatabase.Connection.Test
	correctTestConnectionType := testConnectionType == types.DatabaseSQLiteDatabaseType

	if !correctTestConnectionType {
		return nil, errors.New("test connection type is incorrect")
	}

	return &RelationDatabase{
		MasterORMType:        masterORMType,
		MasterConnectionType: masterConnectionType,
		SlaveORMType:         slaveORMType,
		SlaveConnectionType:  slaveConnectionType,
		TestConnectionType:   testConnectionType,
		Config:               *applicationConfig,
	}, nil
}

func NewFXProvidersRelationDatabase(applicationConfig *config.Config) (fx.Option, error) {
	database, err := NewRelationDatabase(applicationConfig)

	if err != nil {
		return nil, err
	}

	fxProvideRepositories, err := database.GetFXProvideRepositories()

	if err != nil {
		return nil, err
	}

	return fx.Options(
		fxProvideRepositories,
		fx.Provide(repositories.NewMutableRepository),
		fx.Provide(repositories.NewQueryRepository),
	), nil
}

func (database RelationDatabase) GetMasterGormORM() (MasterGormDB, error) {
	var dialector gorm.Dialector

	switch database.MasterConnectionType {
	case types.DatabaseSQLiteDatabaseType:
		dialector = sqlite.Open(database.Config.RelationDatabase.Connections.Sqlite.MasterDSN)
	}

	databaseGorm, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, err
	}

	err = databaseGorm.AutoMigrate(&models.GormPostModel{}, &models.GormUserModel{})

	if err != nil {
		return nil, err
	}

	return databaseGorm, nil
}

func (database RelationDatabase) GetTestGormORM() (TestGormDB, error) {
	var dialector gorm.Dialector

	switch database.TestConnectionType {
	case types.DatabaseSQLiteDatabaseType:
		dialector = sqlite.Open(database.Config.RelationDatabase.Connections.Sqlite.TestDSN)
	}

	databaseGorm, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, err
	}

	err = databaseGorm.AutoMigrate(&models.GormPostModel{}, &models.GormUserModel{})

	if err != nil {
		return nil, err
	}

	return databaseGorm, nil
}

func (database RelationDatabase) GetSlaveGormORM() (SlaveGormDB, error) {
	var dialector gorm.Dialector

	switch database.SlaveConnectionType {
	case types.DatabaseSQLiteDatabaseType:
		dialector = sqlite.Open(database.Config.RelationDatabase.Connections.Sqlite.SlaveDSN)
	}

	databaseGorm, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, err
	}

	return databaseGorm, nil
}

func (database RelationDatabase) GetFXProvideRepositories() (fx.Option, error) {
	var providers []fx.Option

	switch database.MasterORMType {
	case types.DatabaseGormORMType:
		masterDatabase, err := database.GetMasterGormORM()

		if err != nil {
			return nil, err
		}

		providers = append(providers, fx.Options(
			fx.Supply(masterDatabase),
			fx.Provide(
				fx.Annotate(
					func(db MasterGormDB) *repositories.GormPostRepository {
						return repositories.NewGormPostRepository(db)
					},
					fx.As(new(domairepository.MutablePostRepository)),
				),
				fx.Annotate(
					func(db MasterGormDB) *repositories.GormUserRepository {
						return repositories.NewGormUserRepository(db)
					},
					fx.As(new(domairepository.MutableUserRepository)),
				),
			),
		))
	}

	switch database.SlaveORMType {
	case types.DatabaseGormORMType:
		slaveDatabase, err := database.GetSlaveGormORM()

		if err != nil {
			return nil, err
		}

		providers = append(providers, fx.Options(
			fx.Supply(slaveDatabase),
			fx.Provide(
				fx.Annotate(
					func(db SlaveGormDB) *repositories.GormPostRepository {
						return repositories.NewGormPostRepository(db)
					},
					fx.As(new(domairepository.QueryPostRepository)),
				),
				fx.Annotate(
					func(db SlaveGormDB) *repositories.GormUserRepository {
						return repositories.NewGormUserRepository(db)
					},
					fx.As(new(domairepository.QueryUserRepository)),
				),
			),
		))
	}

	return fx.Options(providers...), nil
}
