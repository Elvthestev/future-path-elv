package rest

import (
	"future-path/model"
	"future-path/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) GetUnivNegeri(ctx *gin.Context) {
	namaUniv := ctx.Query("universitas")
	univ, err := r.service.UniversitasService.GetUnivNegeri(namaUniv)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get university", err)
		return
	}
	var univResponse []model.GetUniversitas
	for _, b := range univ {
		univResponse = append(univResponse, model.GetUniversitas{
			Nama_Universitas:      b.Nama_Universitas,
			Alamat_Universitas:    b.Alamat_Universitas,
			Deskripsi_Universitas: b.Deskripsi_Universitas,
		})
	}

	response.Success(ctx, http.StatusOK, "success to get university", univResponse)
}

func (r *Rest) GetUnivSwasta(ctx *gin.Context) {
	namaUniv := ctx.Query("universitas")
	univ, err := r.service.UniversitasService.GetUnivSwasta(namaUniv)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get university", err)
		return
	}
	var univResponse []model.GetUniversitas
	for _, b := range univ {
		univResponse = append(univResponse, model.GetUniversitas{
			Nama_Universitas:      b.Nama_Universitas,
			Alamat_Universitas:    b.Alamat_Universitas,
			Deskripsi_Universitas: b.Deskripsi_Universitas,
		})
	}
	response.Success(ctx, http.StatusOK, "success to get university", univResponse)
}
