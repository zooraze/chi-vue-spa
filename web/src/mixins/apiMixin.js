export default {
  data: function () {
    return {
      apiEndpoint: "/api"
    }
  },

  methods: {
    apiGetData: function () {
      return fetch(`${this.apiEndpoint}/data`)
        .then(resp => {
          return resp.json();
        })
        .catch(err => {
          console.log(`### API Error! ${err}`);
        })
    },
    apiGetList: function () {
      return fetch(`${this.apiEndpoint}/list`)
        .then(resp => {
          return resp.json();
        })
        .catch(err => {
          console.log(`### API Error! ${err}`);
        })
    }
  }
}