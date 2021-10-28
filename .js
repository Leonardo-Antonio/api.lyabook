db.payments.aggregate([
  { $match: { _id: ObjectId("6179001e18a80cce7304bd4b") } },
  {
    $lookup: {
      from: "users",
      localField: "id_client",
      foreignField: "_id",
      as: "client",
    },
  },
])[
  {
    _id: ObjectId("6179001e18a80cce7304bd4b"),
    created_at: ISODate("2021-10-27T07:19:47.179Z"),
    id_client: ObjectId("615cd097adfae7fe7911c491"),
    payment_id: 1242769337,
    status: "approved",
    products: [
      {
        _id: ObjectId("6179001e18a80cce7304bd4c"),
        id_payment: "Kiara Cass",
        title: "La Elegida",
        unit_price: 100,
        quantity: 2,
        description: "f",
        picture_url: "http://localhost:8001/api/v1/images/Q4hE-T_z.png",
        category_id: "-76.9654784 -12.173312",
      },
      {
        _id: ObjectId("6179001e18a80cce7304bd4d"),
        id_payment: "Kiara Cass",
        title: "La Selección",
        unit_price: 50,
        quantity: 1,
        description: "d",
        picture_url: "http://localhost:8001/api/v1/images/c0MLLFon.png",
        category_id: "http://192.168.1.7:8001/api/v1/pdfs/DWPMrDqL.pdf",
      },
      {
        _id: ObjectId("6179001e18a80cce7304bd4e"),
        id_payment: "Charles Dickens",
        title: "Historia De Dos Ciudades",
        unit_price: 100.5,
        quantity: 10,
        description: "f",
        picture_url:
          "http://quelibroleo.com/images/libros/libro_1360114364.jpg",
        category_id: "12515 -1551212",
      },
    ],
    active: true,
    __v: 0,
    client: [
      {
        _id: ObjectId("615cd097adfae7fe7911c491"),
        name: "Alexandra Jaqueline",
        last_name: "Navarro Navarro",
        dni: "72964584",
        password:
          "N9vmZgITzOcSFfDQfr2m3LKN6CEpzGr6LaDZtDxl-QWLXFonDD0HAypIO_8=",
        rol: "Client",
        created_at: ISODate("2021-10-05T22:24:23.894Z"),
        active: true,
      },
    ],
  }
];
