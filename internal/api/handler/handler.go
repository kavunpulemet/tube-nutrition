package handler

import (
	"net/http"
	"strings"
	nutritionservice "tube-nutrition/internal/service"
	"tube-nutrition/internal/utils"
)

func NutritionPlan(ctx utils.MyContext, service nutritionservice.NutritionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		headerParts := strings.Split(authHeader, " ")
		accessToken := headerParts[1]

		nutritionPlan, err := service.NutritionPlan(ctx, accessToken)
		if err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = utils.WriteResponse(w, http.StatusOK, nutritionPlan); err != nil {
			utils.NewErrorResponse(ctx, w, "internal server error", http.StatusInternalServerError)
			return
		}
	}
}
