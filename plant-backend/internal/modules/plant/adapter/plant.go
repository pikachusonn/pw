package adapter

import (
	"context"
	"plant-backend/internal/modules/plant/domain"
)

type PlantRepository struct {
	plants []*domain.Plant
}

func NewPlantRepo() *PlantRepository {
	return &PlantRepository{
		plants: []*domain.Plant{
			domain.NewPlant(
				"1",
				domain.Position{X: 1.0, Y: 2.0, Z: 3.0},
				"Rau ăn",
				"Bị héo la, cần làm thêm chắn nắng",
				[]string{"Rau muống", "Rau mùng tơi"},
				[]string{
					"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT4W5se-3sXcI-CuvSm5GbPoSk655stnvqEeWyX1M79KA&s=10",
					"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT4W5se-3sXcI-CuvSm5GbPoSk655stnvqEeWyX1M79KA&s=10",
				},
			),
			domain.NewPlant(
				"2",
				domain.Position{X: 1.0, Y: 2.0, Z: 3.0},
				"Hoa hồng",
				"Bị sâu",
				[]string{"Hoa hồng xanh", "Hoa hồng đỏ", "Cẩm tú cầu"},
				[]string{
					"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT4W5se-3sXcI-CuvSm5GbPoSk655stnvqEeWyX1M79KA&s=10",
					"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT4W5se-3sXcI-CuvSm5GbPoSk655stnvqEeWyX1M79KA&s=10",
				},
			),
			domain.NewPlant(
				"3",
				domain.Position{X: 1.0, Y: 2.0, Z: 3.0},
				"Chậu cậu tầng 4 - 3",
				"Ngày tưới 1 lần, tưới nhiều úng rễ",
				[]string{"Rau muống", "Rau mùng tơi"},
				[]string{
					"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT4W5se-3sXcI-CuvSm5GbPoSk655stnvqEeWyX1M79KA&s=10",
					"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT4W5se-3sXcI-CuvSm5GbPoSk655stnvqEeWyX1M79KA&s=10",
				},
			),
		},
	}
}

func (r *PlantRepository) ListAll(ctx context.Context) ([]*domain.Plant, error) {
	return r.plants, nil
}
