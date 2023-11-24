import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/RequestView/LoginView.vue'
import RegisterView from "@/views/RequestView/RegisterView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component:()=>import('../views/HomeView/HomeView.vue'),
      children:[
        {
          path:'home_shop_open',
          name:'home_shop_open',
          component:()=>import('../views/HomeView/ShopOpenView.vue')
        },
        {
          path:'home_product_selling',
          name:'home_product_selling',
          component:()=>import('../views/HomeView/ProductSelling.vue')
        },
      ]
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView
    },
    {
      path: '/register',
      name: 'register',
      component: RegisterView
    },
    {
      path:'/sys_user',
      name:'sys_user',
      component: ()=>import('../views/ManagerView/SysUserView.vue'),
      children:[
        {
          path:'all_user_data',
          name:'all_user_data',
          component:()=>import('../views/ManagerView/User/UserData.vue'),
        },
        {
          path:'all_user_account',
          name:'all_user_account',
          component:()=>import('../views/ManagerView/User/UserAccount.vue'),
        },
        {
          path:'all_shop_data',
          name:'all_shop_data',
          component:()=>import('../views/ManagerView/Shop/ShopData.vue'),
        },
        {
          path:'all_shop_account',
          name:'all_shop_account',
          component:()=>import('../views/ManagerView/Shop/ShopAccount.vue'),
        },
        {
          path: 'all_product_data',
          name: 'all_product_data',
          component:()=>import('../views/ManagerView/Product/ProductData.vue')
        },
        {
          path:'shop_request_data',
          name:'shop_request_data',
          component:()=>import('../views/ManagerView/Shop/ShopRequestData.vue'),
        },
        {
          path:'shop_delete_data',
          name:'shop_delete_data',
          component:()=>import('../views/ManagerView/Shop/ShopDelete.vue'),
        },
        {
          path: 'product_request_data',
          name: 'product_request_data',
          component:()=>import('../views/ManagerView/Product/ProductRequest.vue')
        },
        {
          path: 'all_trade_record_data',
          name: 'all_trade_record_data',
          component:()=>import('../views/ManagerView/TradeRecord/AllTradeRecord.vue')
        },
        {
          path: 'not_checked_trade_record_data',
          name: 'not_checked_trade_record_data',
          component:()=>import('../views/ManagerView/TradeRecord/NotCheckedTradeRecord.vue')
        }
      ]
    },
    {
      path:'/user',
      name:'user',
      component: ()=>import('../views/UserView/UserView.vue'),
      children:[
        {
          path:'user_shop_open',
          name:'user_shop_open',
          component:()=>import('../views/HomeView/ShopOpenView.vue')
        },
        {
          path:'user_product_selling',
          name:'user_product_selling',
          component:()=>import('../views/HomeView/ProductSelling.vue')
        },
        {
          path:'user_shop_cart',
          name:'user_shop_cart',
          component:()=>import('../views/UserView/ShopCart.vue')
        },
      ]
    },
    {
      path:'/seller',
      name:'seller',
      component: ()=>import('../views/SellerView/SellerView.vue'),
      children:[
        {
          path:'product_data',
          name:'product_data',
          component:()=>import('../views/SellerView/MyShopView/ProductData.vue'),
        },
        {
          path:'order_data',
          name:'order_data',
          component:()=>import('../views/SellerView/MyShopView/OrderData.vue'),
        },
        {
          path:'trade_record_data',
          name:'trade_record_data',
          component:()=>import('../views/SellerView/MyShopView/TradeRecordData.vue'),
        },
        {
          path:'apply_for_product',
          name:'apply_for_product',
          component:()=>import('../views/RequestView/ApplyForProduct.vue')
        },
        {
          path:'seller_shop_open',
          name:'seller_shop_open',
          component:()=>import('../views/HomeView/ShopOpenView.vue')
        },
        {
          path:'seller_product_selling',
          name:'seller_product_selling',
          component:()=>import('../views/HomeView/ProductSelling.vue')
        },
        {
          path:'seller_shop_cart',
          name:'seller_shop_cart',
          component:()=>import('../views/UserView/ShopCart.vue')
        },
      ]
    },
    {
      path:'/personal_information',
      name:'personal_information',
      component: ()=>import('../views/UserView/UserInformation/PersonLInformation.vue'),
      children:[
        {
          path:'change_username',
          name:'change_username',
          components: {
            username:()=>import('../views/UserView/UserInformation/ChangeUserInformation/ChangeUsername.vue'),
        }},
        {
          path:'change_password',
          name:'change_password',
          components: {
            password:()=>import('../views/UserView/UserInformation/ChangeUserInformation/ChangePassword.vue'),
          },
        },
        {
          path:'change_email',
          name:'change_email',
          components: {
            email:()=>import('../views/UserView/UserInformation/ChangeUserInformation/ChangeEmail.vue'),
          },
        },
        {
          path:'change_phone',
          name:'change_phone',
          components: {
            phone:()=>import('../views/UserView/UserInformation/ChangeUserInformation/ChangePhone.vue')
          },
        },
        {
          path:'recharge_account',
          name:'recharge_account',
          components: {
            balance:()=>import('../views/UserView/UserInformation/RechargeAccount.vue'),
          }},
      ]
    },
    {
      path:'/apply_for_shop',
      name:'apply_for_shop',
      component: ()=>import('../views/RequestView/ApplyForShop.vue')
    },
    {
      path:'/seller_shop_information',
      name:'seller_shop_information',
      component: ()=>import('../views/SellerView/MyShopView/ShopInformation.vue'),
      children:[
        {
          path:'recharge_shop_account',
          name:'recharge_shop_account',
          components: {
            shop_balance:()=>import('../views/SellerView/MyShopView/RechargeShopAccount.vue'),
          }},
      ]
    },
  ]
})

export default router
