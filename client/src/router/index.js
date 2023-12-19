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
          path:'/home_item_all',
          name:'home_item_all',
          component:()=>import('../views/HomeView/CommodityItemAll.vue')
        },
        {
          path:'/home_commodity_all',
          name:'home_commodity_all',
          component:()=>import('../views/UserView/CommodityAll.vue')
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
          path: 'all_admin_commodity_item_data',
          name: 'all_admin_commodity_item_data',
          component:()=>import('../views/AdminView/Commodity/CommodityItemData.vue')
        },
        {
          path: 'all_platform_data',
          name: 'all_platform_data',
          component:()=>import('../views/AdminView/Platform/PlatformData.vue')
        },
        {
          path: 'statistics_favorite',
          name: 'statistics_favorite',
          component:()=>import('../views/AdminView/Statistics/StatisticsFavorite.vue')
        },
        {
          path: 'statistics_commodity_price',
          name: 'statistics_commodity_price',
          component:()=>import('../views/AdminView/Statistics/StatisticsCommodityPrice.vue')
        },
        {
          path: 'statistics_annual',
          name: 'statistics_annual',
          component:()=>import('../views/AdminView/Statistics/StatisticsAnnual.vue')
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
          path:'user_commodity_data',
          name:'user_commodity_data',
          component:()=>import('../views/UserView/CommodityAll.vue')
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
        {
          path:'search_commodity',
          name:'search_commodity',
          component:()=>import('../views/UserView/SearchCommodity.vue')
        },
        {
          path:'user_annual',
          name:'user_annual',
          component:()=>import('../views/UserView/UserAnnual.vue')
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
