<template>
  <el-container>
    <el-header style="background-color: darkgrey;border-radius: 5px" height="65px">
      <el-container>
        <el-aside
          width="500px"
          style="text-align: left; font-size: 20px; font-weight: 1000; color: darkslateblue"
        >
          <div style="margin-top: 15px">
            <span>PriceCompareSystem</span>
          </div>
        </el-aside>
        <el-main style="text-align: right; font-size: 16px">
          <div>
            <span>未登录</span>
            <el-dropdown>
              <el-icon style="margin-left: 8px; margin-top: 1px; font-size: large">
                <setting />
              </el-icon>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item>
                    <RouterLink to="/login">登录</RouterLink>
                  </el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </div>
        </el-main>
      </el-container>
    </el-header>
    <el-container>
      <el-aside width="200px" height="640px">
        <el-container>
          <el-header
            class="layout_header"
            style="text-align: center; font-size: 18px; background-color: antiquewhite;border-radius: 5px;margin-right: 20px"
            height="50px"
          >
            <div style="margin-top: 10px;">
              <span>Navigation</span>
            </div>
          </el-header>
          <el-main>
            <el-menu>
              <el-menu-item index="1">
                <template #title>
                  <el-icon>
                    <goods />
                  </el-icon>
                  <div @click="router.push('/home_item_all')">所有商品实例</div>
                </template>
              </el-menu-item>
              <el-menu-item index="2">
                <template #title>
                  <el-icon>
                    <goods-filled />
                  </el-icon>
                  <div @click="router.push('/home_commodity_all')">所有商品</div>
                </template>
              </el-menu-item>
              <el-menu-item index="3">
                <template #title>
                  <el-icon>
                    <user />
                  </el-icon>
                  <div @click="remind_to_login">个人信息</div>
                </template>
              </el-menu-item>
              <el-menu-item index="4">
                <template #title>
                  <el-icon>
                    <star />
                  </el-icon>
                  <div @click="remind_to_login">个人收藏</div>
                </template>
              </el-menu-item>
              <el-menu-item index="5">
                <template #title>
                  <el-icon>
                    <message />
                  </el-icon>
                  <div @click="remind_to_login">系统消息</div>
                </template>
              </el-menu-item>
            </el-menu>
          </el-main>
        </el-container>
      </el-aside>
      <el-main height="605px">
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import { ElMessage } from "element-plus";

export default {
  name: "home_view",
  data() {
    return {
    };
  },
  created() {
  },
  methods: {
  }
};

const remind_to_login = () => {
  ElMessage.error('请先登录！')
}
</script>

<script setup>
// eslint-disable-next-line no-unused-vars
import { GoodsFilled, Menu as IconMenu, Message, Setting } from "@element-plus/icons-vue";
import { RouterLink } from "vue-router";
import router from "@/router";

</script>

