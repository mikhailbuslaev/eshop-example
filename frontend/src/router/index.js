import { createRouter, createWebHashHistory } from 'vue-router'
import CatalogView from '../views/CatalogView.vue'
import AboutView from '../views/AboutView.vue'
import CardView from '../views/CardView.vue'
import ShoppingCartView from '../views/ShoppingCartView.vue'
import AuthorizationView from '../views/AuthorizationView.vue'
import OrdersView from '../views/OrdersView.vue'
const routes = [
  {
    path: '/catalog/:category/:page',
    name: 'catalog',
    component: CatalogView
  },
  {
    path: '/about',
    name: 'about',
    component: AboutView
  },
  { 
    path: '/product/:category/:id', 
    name: 'card',
    component: CardView 
  },
  { 
    path: '/user/shopping_cart', 
    name: 'shopping_cart',
    component: ShoppingCartView
  },
  {
    path: '/login', 
    name: 'login',
    component: AuthorizationView
  },
  {
    path: '/user/orders/:page',
    name: 'orders',
    component: OrdersView
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
