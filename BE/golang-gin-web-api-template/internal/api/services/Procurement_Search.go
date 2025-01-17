package services

import (
	"fmt"

	"web-api/internal/pkg/config"
	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/request"
	"web-api/internal/pkg/models/types"
)

type ProcService struct {
	*BaseService
}

var Proc = &ProcService{}

func (s *ProcService) GetPurList(requestParams request.ProcListRequest) ([]types.ProcList, error) {
	configuration := config.GetConfig()
	configuration.Database.Driver = "sqlserver"
	configuration.Database.Host = "192.168.23.122"
	configuration.Database.Username = "tyxuan"
	configuration.Database.Password = "Erp@admin2309"
	configuration.Database.Dbname = "LIY_ERP"
	configuration.Database.Port = "1433"
	db, err := database.CreateDatabaseConnection(configuration)
	if err != nil {
		return []types.ProcList{}, err
	}
	dbInstance, _ := db.DB()

	// Sử dụng fmt.Sprintf để chèn giá trị động vào câu lệnh SQL
	query := fmt.Sprintf(`select *, DATEDIFF(DAY, Startdate, EndDate) as DiffDay from
	(
		select DDBH,
		(select top 1 USERDATE from KCRKS where CGBH = ddzl.DDBH order by USERDATE) as StartDate,
		(select top 1 INDATE from YWCP where DDBH = ddzl.DDBH order by INDATE Desc) as EndDate
		from (
			select *,(select sum(QTY) from YWCP
			where DDBH = ddzl.DDBH) as OKQTY from (
				select DDBH, Pairs from ddzl
				where left(BUYNO, 6) = '%s' and GSBH = '%s' and pairs > 1
			) ddzl
		) ddzl
		where pairs = OKQTY
	) YWCP1
	order by DiffDay`, requestParams.BUYNO, requestParams.GSBH)
	var result []types.ProcList
	err = db.Raw(query).Scan(&result).Error
	dbInstance.Close()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *ProcService) GetAverageDiffDay(requestParams request.ProcListRequest) (float64, error) {
	configuration := config.GetConfig()
	configuration.Database.Driver = "sqlserver"
	configuration.Database.Host = "192.168.23.122"
	configuration.Database.Username = "tyxuan"
	configuration.Database.Password = "Erp@admin2309"
	configuration.Database.Dbname = "LIY_ERP"
	configuration.Database.Port = "1433"
	db, err := database.CreateDatabaseConnection(configuration)
	if err != nil {
		return 0, err
	}
	dbInstance, _ := db.DB()

	// Sử dụng fmt.Sprintf để chèn giá trị động vào câu lệnh SQL
	query := fmt.Sprintf(`
		SELECT CAST(AVG(CAST(s.DiffDay AS FLOAT)) AS DECIMAL(10, 4)) AS AverageDiffDay
		FROM (
			SELECT *, DATEDIFF(DAY, Startdate, EndDate) AS DiffDay
			FROM (
				SELECT DDBH,
				       (SELECT TOP 1 USERDATE FROM KCRKS WHERE CGBH = ddzl.DDBH ORDER BY USERDATE) AS StartDate,
				       (SELECT TOP 1 INDATE FROM YWCP WHERE DDBH = ddzl.DDBH ORDER BY INDATE DESC) AS EndDate
				FROM (
					SELECT *, (SELECT SUM(QTY) FROM YWCP WHERE DDBH = ddzl.DDBH) AS OKQTY
					FROM (
						SELECT DDBH, Pairs
						FROM ddzl
						WHERE LEFT(BUYNO, 6) = '%s' AND GSBH = '%s' AND Pairs > 1
					) ddzl
				) ddzl
				WHERE Pairs = OKQTY
			) YWCP1
		) AS s
	`, requestParams.BUYNO, requestParams.GSBH)

	var averageDiffDay float64
	err = db.Raw(query).Scan(&averageDiffDay).Error
	dbInstance.Close()
	if err != nil {
		return 0, err
	}

	return averageDiffDay, nil
}
