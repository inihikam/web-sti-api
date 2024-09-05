package services

import (
	"github.com/inihikam/web-sti-api/database"
	"github.com/inihikam/web-sti-api/models"
)  

func GetLecturers() ([]models.LecturersByPosition, error) {  
    err := database.InitDB()  
    if err != nil {  
        return nil, err  
    }  
    defer database.DB.Close()  

    query := `  
        SELECT  
            d.id, d.nama, d.npp,  
            j.id, j.nama  
        FROM dosen d  
        JOIN jabatan j ON d.jabatan_id = j.id  
    `  
    rows, err := database.DB.Query(query)  
    if err != nil {  
        return nil, err  
    }  
    defer rows.Close()  

    var lecturersByPosition []models.LecturersByPosition  
    lecturerMap := make(map[string][]models.Dosen)  
    for rows.Next() {  
        var lecturer models.Dosen  
        var jabatan models.Jabatan  
        if err := rows.Scan(&lecturer.ID, &lecturer.Nama, &lecturer.NPP, &jabatan.ID, &jabatan.Nama); err != nil {  
            return nil, err  
        }  
        lecturer.Jabatan = jabatan  
        lecturerMap[jabatan.Nama] = append(lecturerMap[jabatan.Nama], lecturer)  
    }  

    if err := rows.Err(); err != nil {  
        return nil, err  
    }  

    for position, lecturers := range lecturerMap {  
        lecturersByPosition = append(lecturersByPosition, models.LecturersByPosition{  
            Position:  position,  
            Lecturers: lecturers,  
        })  
    }  

    return lecturersByPosition, nil  
}