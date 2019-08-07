import Vue from 'vue'
import Router from 'vue-router'
import Home from './components/Home.vue'
import Data from './components/Data.vue'
import Error from './components/Error.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/home',
      name: 'apphome',
      component: Home
    },
    {
      path: '/data',
      name: 'data',
      component: Data
    },
    {
      path: '*',
      name: 'catchall',
      component: Error
    }
  ]
})
