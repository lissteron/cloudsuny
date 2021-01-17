function GetItems(axios, params) {
  if (!params.fields) {
    params.fields = ['id', 'url', 'name', 'price', 'price_name', 'min_amount', 'discount_size', 'photos', 'show'];
  }

  return new Promise((resolve, reject) => {
    axios
      .post('/api/v1/list_items', params)
      .then((res) => {
        resolve(res.data.data);
      })
      .catch((e) => {
        reject(e);
      });
  });
}

export default { GetItems };
