// Vue.js entry point
import Vue from 'vue'
import App from './App.vue'
import router from './router'

// UI and Bootstrap stuff
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import './scss/theme.scss'

// Import Font Awesome icons
import { library as fontAwesomeLib } from '@fortawesome/fontawesome-svg-core'
import { faHome, faCogs, faTachometerAlt, faInfoCircle, faUmbrella, faBomb } from '@fortawesome/free-solid-svg-icons'
import { faGithub, faDocker } from '@fortawesome/free-brands-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
// Register icons and component
fontAwesomeLib.add([faHome, faCogs, faTachometerAlt, faInfoCircle, faGithub, faDocker, faUmbrella, faBomb])
Vue.component('fa', FontAwesomeIcon)

// Register globally
import VueSkycons from 'vue-skycons'
Vue.use(VueSkycons)

// Init Vue 
Vue.use(BootstrapVue);
Vue.config.productionTip = false

// Root Vue instance
// Mount on the <div id="root"> and render the template of the App component
new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
