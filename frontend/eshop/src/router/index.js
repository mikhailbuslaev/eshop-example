import { createRouter, createWebHashHistory } from 'vue-router'
import CatalogView from '../views/CatalogView.vue'
import AboutView from '../views/AboutView.vue'
import CardView from '../views/CardView.vue'
const routes = [
  {
    path: '/catalog/:page',
    name: 'catalog',
    component: CatalogView
  },
  {
    path: '/about',
    name: 'about',
    component: AboutView
  },
  { 
    path: '/product/:id', 
    name: 'card',
    component: CardView 
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
