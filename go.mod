module boulihub

go 1.24.3

replace boulihub/utils => ./utils

replace boulihub/models => ./models

require (
	boulihub/models v0.0.0-00010101000000-000000000000
	boulihub/routes v0.0.0-00010101000000-000000000000
	boulihub/utils v0.0.0-00010101000000-000000000000
)

replace boulihub/routes => ./routes
