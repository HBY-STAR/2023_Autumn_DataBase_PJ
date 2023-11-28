import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    //主界面
    {
      path: '/',
      name: 'home',
      component:()=>import('../views/HomeView/HomeView.vue'),
      children: [
        {
          path:'home_commodity_all',
          name:'home_commodity_all',
          component:()=>import('../views/HomeView/CommodityItemAll.vue')
        },
        {
          path:'search',
          name:'search',
          component:()=>import('../views/HomeView/SearchView.vue')
        },
      ]
    },

    //登录界面
    {
      path: '/login',
      name: 'login',
      component:()=>import('../views/RequestView/LoginView.vue')
    },
    //管理员界面
    {
      path:'/admin',
      name:'admin',
      component: ()=>import('../views/AdminView/AdminView.vue'),
      children:[
        //Get all
        {
          path:'all_user_data',
          name:'all_user_data',
          component:()=>import('../views/AdminView/User/UserData.vue'),
        },
        {
          path:'all_seller_data',
          name:'all_seller_data',
          component:()=>import('../views/AdminView/Seller/SellerData.vue'),
        },
        {
          path: 'all_commodity_data',
          name: 'all_commodity_data',
          component:()=>import('../views/AdminView/Commodity/CommodityData.vue')
        },
        {
          path: 'all_commodity_item_data',
          name: 'all_commodity_item_data',
          component:()=>import('../views/AdminView/Commodity/CommodityItemData.vue')
        },
        {
          path: 'all_platform_data',
          name: 'all_platform_data',
          component:()=>import('../views/AdminView/Platform/PlatformData.vue')
        },
        //Add one
        {
          path:'user_add',
          name:'user_add_data',
          component:()=>import('../views/AdminView/User/UserAdd.vue'),
        },
        {
          path:'seller_add',
          name:'seller_add_data',
          component:()=>import('../views/AdminView/Seller/SellerAddData.vue'),
        },
        {
          path: 'commodity_add',
          name: 'commodity_add',
          component:()=>import('../views/AdminView/Commodity/CommodityAdd.vue')
        },
        {
          path: 'commodity_item_add',
          name: 'commodity_item_add',
          component:()=>import('../views/AdminView/Commodity/CommodityItemAdd.vue')
        },
        {
          path:'platform_add',
          name:'platform_add',
          component:()=>import('../views/AdminView/Platform/PlatformAdd.vue'),
        },
      ]
    },
    //用户界面
    {
      path:'/user',
      name:'user',
      component: ()=>import('../views/UserView/UserView.vue'),
      children:[
        {
          path:'user_data',
          name:'user_data',
          component: ()=>import('../views/UserView/UserData.vue')
        },
        {
          path:'all_commodity_item_data',
          name:'all_commodity_item_data',
          component:()=>import('../views/HomeView/CommodityItemAll.vue')
        },
        {
          path:'favorite',
          name:'favorite',
          component:()=>import('../views/UserView/Favorite.vue')
        },
        {
          path: 'message',
          name: 'message',
          component:() =>import('../views/UserView/MessageView.vue')
        },
        {
          path:'search',
          name:'search',
          component:()=>import('../views/HomeView/SearchView.vue')
        },
      ]
    },
    //商家界面
    {
      path:'/seller',
      name:'seller',
      component: ()=>import('../views/SellerView/SellerView.vue'),
      children:[
        {
          path:'seller_data',
          name:'seller_data',
          component: ()=>import('../views/SellerView/SellerData.vue')
        },
        {
          path:'commodity_item_data',
          name:'commodity_item_data',
          component:()=>import('../views/SellerView/CommodityItemData.vue'),
        },
        {
          path:'add_commodity_item',
          name:'add_commodity_item',
          component:()=>import('../views/SellerView/CommodityItemAdd.vue')
        },
      ]
    },
    //测试界面
    {
      path: '/test',
      name: 'test',
      component:()=>import('../views/TestView.vue')
    },
  ]
})

export default router
