import Vue from 'vue'
import Router from 'vue-router'
import TopPage from '@/components/TopPage'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'TopPage',
      component: TopPage
    },
    {
      path: '/papers/:arxiv_id',
      name: 'TopPage',
      component: TopPage
    },
  ]
})
