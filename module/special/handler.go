package special

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
)

func Create(c echo.Context) error {
	st := new(Specialteam)
	c.Bind(st)
	contentType := c.Request().Header.Get("Content-type")
	if contentType == "application/json" {
		fmt.Println("Request dari json")
	} else if strings.Contains(contentType, "multipart/form-data") || contentType == "application/x-www-form-urlencoded" {
		file, err := c.FormFile("image")
		if err != nil {
			fmt.Println("NO Image")
		} else {
			src, err := file.Open()
			if err != nil {
				return err
			}
			defer src.Close()
			//setname
			t := time.Now()
			uniqueFN := "storage/" + t.Format("20060102150405") + "-" + file.Filename

			dst, err := os.Create(uniqueFN)
			fmt.Println(dst)
			if err != nil {
				return err
			}
			defer dst.Close()
			if _, err = io.Copy(dst, src); err != nil {
				return err
			}

			st.Image = uniqueFN
			fmt.Println("Ada file, akan disimpan")
		}
	}

	response := new(Response)
	if st.CreateSpecialteam() != nil { // method create user
		response.ErrorCode = 10
		response.Message = "Gagal create data user"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses create data user"
		response.Data = *st
	}
	return c.JSON(http.StatusOK, response)
}

func Update(c echo.Context) error {
	st := new(Specialteam)
	c.Bind(st)
	response := new(Response)

	//covert id
	param := c.Param("id")
	paramConverted, _ := strconv.Atoi(param)

	if st.UpdateSpecialteam(paramConverted) != nil { // method update
		response.ErrorCode = 10
		response.Message = "Gagal update data specialTeam"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses update data specialTeam"
		response.Data = *st
	}
	return c.JSON(http.StatusOK, response)
}

func Delete(c echo.Context) error {
	//covert id
	param := c.Param("id")
	paramConverted, _ := strconv.Atoi(param)

	specialTeam, _ := GetOneById(paramConverted) // method get by id
	response := new(Response)

	if specialTeam.DeleteSpecialteam() != nil { // method update specialTeam
		response.ErrorCode = 10
		response.Message = "Gagal menghapus data specialTeam"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses menghapus data specialTeam"
	}
	return c.JSON(http.StatusOK, response)
}

func Search(c echo.Context) error {
	response := new(Response)
	specialTeams, err := GetAll(c.QueryParam("keywords")) // method get all
	if err != nil {
		response.ErrorCode = 10
		response.Message = "Gagal melihat data specialTeam"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses melihat data specialTeam"
		response.Data = specialTeams
	}
	return c.JSON(http.StatusOK, response)
}
