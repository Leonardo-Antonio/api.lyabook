db.payments.aggregate(
    [
        {
            $match: {
                _id: ObjectId("6167d3817d2fd42d6f5fdd7b")
            }
        },
        {
            $lookup: {
                from: 'users',
                localField: 'id_client',
                foreignField: '_id',
                as: 'client'
            }
        }
    ]
)