<template>
  <div
      style="
        margin-top: 140px;
        margin-left: 350px;
        width: 700px;
        height: 460px;
        padding: 20px;
        border-radius: 10px;
        background-color: #909399;
      "
  >
    <div style="width: 300px; float: left">
      <el-card style="width: 300px;">
        <el-form label-width="80px" size="small">
          <div style="height: 30px">
            <header style="text-align: center; color: #337ecc">
              个人信息
            </header>
          </div>
          <el-form-item label="用户序列号:">
            <span>{{ user_data.uid }}</span>
          </el-form-item>
          <el-form-item label="用户名:">
          <span>
            {{ user_data.username }}
            <RouterLink style="position: absolute;right: 0;" to="/personal_information/change_username">
            <el-icon size="medium">
              <edit>
              </edit>
            </el-icon>
          </RouterLink>
          </span>
            <RouterView name="username"></RouterView>
          </el-form-item>
          <el-form-item label="密码:">
          <span>
            {{ user_data.password }}
            <RouterLink style="position: absolute;right: 0;" to="/personal_information/change_password">
            <el-icon size="medium">
              <edit>
              </edit>
            </el-icon>
          </RouterLink>
          </span>
            <RouterView name="password"></RouterView>
          </el-form-item>
          <el-form-item label="邮箱:">
          <span>
            {{ user_data.email }}
            <RouterLink style="position: absolute;right: 0;" to="/personal_information/change_email">
            <el-icon size="medium">
              <edit>
              </edit>
            </el-icon>
          </RouterLink>
          </span>
            <RouterView name="email"></RouterView>
          </el-form-item>
          <el-form-item label="电话:">
          <span>
            {{ user_data.phone }}
            <RouterLink style="position: absolute;right: 0;" to="/personal_information/change_phone">
            <el-icon size="medium">
              <edit>
              </edit>
            </el-icon>
          </RouterLink>
          </span>
            <RouterView name="phone"></RouterView>
          </el-form-item>
          <el-form-item label="用户类型:">
            <span>{{ user_data.userType }}</span>
          </el-form-item>
          <el-form-item label="身份证号:">
            <span>{{ user_data.idNumber }}</span>
          </el-form-item>
        </el-form>
      </el-card>
    </div>
    <div style="width: 100px; height: 100px; float: left">
    </div>
    <div style="width: 300px; float: left">
      <el-card style="width: 300px;">
        <el-form label-width="80px" size="small">
          <div style="height: 30px">
            <header style="text-align: center; color: #337ecc">
              账户信息
            </header>
          </div>
          <el-form-item label="账户ID:">
            <span>{{ user_account.uid }}</span>
          </el-form-item>
          <el-form-item label="账户状态:">
            <span>{{ user_account.accountState }}</span>
          </el-form-item>
          <el-form-item label="账户余额:">
            <div style="width: 300px">
            <span>
              {{ user_account.balance }}
              <RouterLink style="position: absolute;right: 0;" to="/personal_information/recharge_account">
                <el-icon size="medium">
                  <el-tooltip content="充值" effect="light" placement="right">
                    <money style="outline-width: 0">
                    </money>
                  </el-tooltip>
                </el-icon>
              </RouterLink>
            </span>
            </div>
            <RouterView name="balance"></RouterView>
          </el-form-item>
        </el-form>
      </el-card>
      <el-form-item style="margin-top: 140px">
        <el-button
            autocomplete="off"
            style="position: absolute;right: 0; margin-top: 30px"
            type="warning"
            @click="returnToMain"
        >返回主界面
        </el-button
        >
      </el-form-item>
    </div>
  </div>
</template>

<script>
import {Edit, Money} from "@element-plus/icons-vue";

export default {
  name: "personal_information",
  components: {Money, Edit},
  data() {
    return {
      user_data: localStorage.getItem('user_data')
          ? JSON.parse(localStorage.getItem('user_data'))
          : {},
      user_account: localStorage.getItem('user_account')
          ? JSON.parse(localStorage.getItem('user_account'))
          : {
            uid: 0,
            accountState: '',
            balance: 0.0,
            shop_cart: '',
          }
    }
  },
  created() {
    this.findAccount()
  },
  methods: {
    findAccount() {
      this.request.get('/user-account/find_' + this.user_data.uid).then((res) => {
        if (res.code === '200') {
          localStorage.setItem('user_account', JSON.stringify(res.data)) // 存储用户信息到浏览器
        } else {
          this.$message.error(res.msg)
        }
      })
    },
    returnToMain(){
      if (this.user_data.userType === 'sys_user') {
        this.$router.push('/sys_user')
      } else if (this.user_data.userType === 'user') {
        this.$router.push('/user')
      } else if (this.user_data.userType === 'seller') {
        this.$router.push('/seller')
      } else {
        this.$router.push('/')
      }
    }
  }
}
</script>

