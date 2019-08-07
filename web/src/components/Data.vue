<template>
  <div>
    <div class="row justify-content-center">
      <h5 class="p-3 text-center">Arbitrary data for your viewing pleasure...</h5>

      <spinner v-if="!data"></spinner>

      <table v-if="data" class="table table-hover">
        <tbody>
          <tr v-for="(val, key) in dataComputed" :key="key">
            <td>
              <b>{{ key | titleify }}</b>
            </td>
            <td>{{ val }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import apiMixin from "../mixins/apiMixin.js";
import Spinner from "./Spinner.vue";
const data = null;

export default {
  mixins: [apiMixin],

  data: function() {
    return {
      data: data
    };
  },

  components: {
    Spinner
  },

  created() {
    this.getData();
    setInterval(this.getData, 5000);
  },

  methods: {
    getData: function() {
      fetch(`${this.apiEndpoint}/data`)
        .then(resp => {
          return resp.json();
        })
        .then(json => {
          this.data = json;
        })
        .catch(err => {
          console.log(err);
        });
    }
  },

  filters: {
    titleify: function(value) {
      if (!value) return "";
      value = value.toString();
      value = value.replace(/([A-Z])/g, " $1");
      value = value.replace(/^./, function(str) {
        return str.toUpperCase();
      });
      return value;
    }
  },

  computed: {
    dataComputed: function() {
      var result = {};
      for (let k in this.data) {
        result[k] = this.data[k];
      }
      return result;
    }
  }
};
</script>

<style scoped>
pre {
  background-color: #222;
  color: rgb(59, 190, 33);
  padding: 10px;
  max-height: 500px;
  font-family: "Lucida Console", monospace;
}
</style>