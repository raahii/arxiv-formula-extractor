import Vue from 'vue'
import Router from 'vue-router'
import TopPage from '@/components/TopPage'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'TopPage',
      component: TopPage
    },
  ]
})
