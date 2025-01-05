package stats

import (
	"net/http"
	"sportin/config"
	"sportin/database/dbmodel"
	"sportin/pkg/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type StatsConfigurator struct {
	*config.Config
}

func New(configuration *config.Config) *StatsConfigurator {
	return &StatsConfigurator{configuration}
}

func statsToModel(allStats []*dbmodel.Stats) []models.UserStats {
	statsToModel := &models.UserStats{}
	statsEdited := []models.UserStats{}
	for _, stats := range allStats {
		statsToModel.UserId = stats.UserId
		statsToModel.Weight = stats.Weight
		statsToModel.Height = stats.Height
		statsToModel.Age = stats.Age
		statsToModel.ActivityCoefficient = stats.ActivityCoefficient
		statsToModel.CaloriesGoal = stats.CaloriesGoal
		statsToModel.ProteinRatio = stats.ProteinRatio
		statsEdited = append(statsEdited, *statsToModel)
	}
	return statsEdited
}

func (config *StatsConfigurator) statsHandler(w http.ResponseWriter, r *http.Request) {
	stats, err := config.StatsRepository.FindAll()
	statsEdited := statsToModel(stats)
	if err != nil {
		render.JSON(w, r, map[string]string{"Error": "Failed to load all the stats"})
		return
	}
	render.JSON(w, r, statsEdited)
}

func (config *StatsConfigurator) statByIdHandler(w http.ResponseWriter, r *http.Request) {
	statsId := chi.URLParam(r, "id")
	stats, err := config.StatsRepository.FindById(statsId)
	statsEdited := statsToModel(stats)
	if err != nil {
		render.JSON(w, r, map[string]string{"Error": "Failed to load the wanted stat"})
		return
	}
	render.JSON(w, r, statsEdited)
}

func (config *StatsConfigurator) addStatsHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.UserStats{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	addStats := &dbmodel.Stats{UserId: req.UserId, Weight: req.Weight, Height: req.Height, Age: req.Age, ActivityCoefficient: req.ActivityCoefficient, CaloriesGoal: req.CaloriesGoal, ProteinRatio: req.ProteinRatio}
	config.StatsRepository.Create(addStats)
	render.JSON(w, r, map[string]string{"success": "New stats successfully added"})
}

func (config *StatsConfigurator) editStatsHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.UserStats{}
	statsId := chi.URLParam(r, "id")
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	updatedStats := &dbmodel.Stats{UserId: req.UserId, Weight: req.Weight, Height: req.Height, Age: req.Age, ActivityCoefficient: req.ActivityCoefficient, CaloriesGoal: req.CaloriesGoal, ProteinRatio: req.ProteinRatio}
	config.StatsRepository.Update(updatedStats, statsId)
	render.JSON(w, r, map[string]string{"success": "Stats successfully updated"})
}

func (config *StatsConfigurator) deleteStatsHandler(w http.ResponseWriter, r *http.Request) {
	statsId := chi.URLParam(r, "id")
	stats, err := config.StatsRepository.FindById(statsId)
	if err != nil {
		render.JSON(w, r, map[string]string{"Error": "Failed to find the wanted stats"})
		return
	}
	config.StatsRepository.Delete(stats[0])
	render.JSON(w, r, map[string]string{"success": "Stats successfully deleted"})
}
