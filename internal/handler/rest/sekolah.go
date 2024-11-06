package rest

import (
	"future-path/model"
	"future-path/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) GetSekolahNegeri(ctx *gin.Context) {
	namaQuery := ctx.Query("sekolah")
	sekolah, err := r.service.SekolahService.GetSekolahNegeri(namaQuery)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get school", err)
		return
	}

	var sekolahResponses []model.GetSekolah
	for _, b := range sekolah {
		sekolahResponses = append(sekolahResponses, model.GetSekolah{
			Nama_Sekolah:      b.Nama_Sekolah,
			Alamat_Sekolah:    b.Alamat_Sekolah,
			Deskripsi_Sekolah: b.Alamat_Sekolah,
		})
	}

	response.Success(ctx, http.StatusOK, "success to get school", sekolahResponses)
}

func (r *Rest) GetSekolahSwasta(ctx *gin.Context) {
	namaQuery := ctx.Query("sekolah")
	sekolah, err := r.service.SekolahService.GetSekolahSwasta(namaQuery)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get school", err)
		return
	}

	var sekolahResponses []model.GetSekolah
	for _, b := range sekolah {
		sekolahResponses = append(sekolahResponses, model.GetSekolah{
			Nama_Sekolah:      b.Nama_Sekolah,
			Alamat_Sekolah:    b.Alamat_Sekolah,
			Deskripsi_Sekolah: b.Alamat_Sekolah,
		})
	}

	response.Success(ctx, http.StatusOK, "success to get school", sekolahResponses)
}
