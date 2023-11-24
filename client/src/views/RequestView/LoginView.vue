<template>
  <div class="wrapper">
    <div
      style="
        margin: 200px auto;
        background-color: #fff;
        width: 300px;
        height: 250px;
        padding: 20px;
        border-radius: 10px;
      "
    >
      <div style="margin: 20px 0; text-align: center; font-size: 24px"><b>登录</b></div>
      <el-form :model="user" :rules="rules" ref="userForm">
        <el-form-item prop="username">
          <el-input
            size="large"
            prefix-icon="User"
            v-model="user.username"
            placeholder="请输入用户名"
          ></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            size="large"
            prefix-icon="Lock"
            show-password
            v-model="user.password"
            placeholder="请输入密码"
          ></el-input>
        </el-form-item>
        <el-form-item style="margin: 40px 0; text-align: right">
          <el-space direction="horizontal" alignment="center" :size="180">
            <el-button type="primary" size="default" autocomplete="off" @click="login"
              >登录</el-button
            >
            <el-button
              type="warning"
              size="default"
              autocomplete="off"
              @click="$router.push('/register')"
              >注册</el-button
            >
          </el-space>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'user_login',
  data() {
    return {
      user: {},
      rules: {
        username: [
                { required: true, message: '用户名不能为空', trigger: 'blur' },
          { min: 3, max: 10, message: '长度在 3 到 10 个字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '密码不能为空', trigger: 'blur' },
          { min: 6, max: 32, message: '长度在 6 到 32 个字符', trigger: 'blur' }
        ]
      },
      user_data: {
        uid: 0,
        username: '',
        password: '',
        email: '',
        phone: '',
        userType: '',
        idNumber: '',
        avatarUrl:'',
      }
    }
  },
  methods: {
    login() {
      this.$refs['userForm'].validate((valid) => {
        if (valid) {
          // 表单校验合法
          //axios.post("/user-data/login", this.user).then()
          this.request.post('/user-data/login', this.user).then((res) => {
            if (res.code === '200') {
              localStorage.setItem('user_data', JSON.stringify(res.data)) // 存储用户信息到浏览器
              if (res.data.userType === 'sys_user') {
                this.$router.push('/sys_user')
              } else if (res.data.userType === 'user') {
                this.$router.push('/user')
              } else if (res.data.userType === 'seller') {
                this.$router.push('/seller')
              } else {
                this.$router.push('/')
              }
              this.$message.success('登录成功')
            } else {
              this.$message.error(res.msg)
            }
          })
        }
      })
    }
  }
}
</script>

<style scoped>
.wrapper {
  height: 100vh;
  background-color: #909399;
  overflow: hidden;
}
</style>
