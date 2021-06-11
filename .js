db.student_teacher.aggregate(
    [
        {
            $match: {
                id_teacher: ObjectId("60c2b583ba5547a52c98f346"),
                active: true
            }
        },
        {
            $lookup: {
                from: 'users',
                localField: 'id_client',
                foreignField: '_id',
                as: 'students', 
            }
        }
    ]
).pretty()