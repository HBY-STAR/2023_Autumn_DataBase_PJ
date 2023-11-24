<template>
  <div class="wrapper">
    <div
      style="
        margin: 100px auto;
        background-color: #fff;
        width: 350px;
        height: 480px;
        padding: 20px;
        border-radius: 10px;
      "
    >
      <div style="margin: 20px 0; text-align: center; font-size: 24px"><b>注 册</b></div>
      <el-form :model="user" :rules="rules" ref="userForm">
        <el-form-item prop="email">
          <el-input
            placeholder="请输入邮箱"
            size="medium"
            prefix-icon="Message"
            v-model="user.email"
          ></el-input>
        </el-form-item>
        <el-form-item prop="phone">
          <el-input
            placeholder="请输入手机号"
            size="medium"
            prefix-icon="Iphone"
            v-model="user.phone"
          ></el-input>
        </el-form-item>
        <el-form-item prop="id_number">
          <el-input
            placeholder="请输入身份证号"
            size="medium"
            prefix-icon="WarningFilled"
            v-model="user.id_number"
          ></el-input>
        </el-form-item>
        <el-form-item prop="username">
          <el-input
            placeholder="请输入用户名"
            size="medium"
            prefix-icon="User"
            v-model="user.username"
          ></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            placeholder="请输入密码"
            size="medium"
            prefix-icon="Lock"
            show-password
            v-model="user.password"
          ></el-input>
        </el-form-item>
        <el-form-item prop="confirmPassword">
          <el-input
            placeholder="请确认密码"
            size="medium"
            prefix-icon="Lock"
            show-password
            v-model="user.confirmPassword"
          ></el-input>
        </el-form-item>
        <el-form-item label="请选择用户类型" prop="user_type">
          <el-radio-group v-model="user.user_type">
            <el-radio border label="user">普通用户</el-radio>
            <el-radio border label="seller">商户</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item style="margin: 30px 0; text-align: right">
          <el-space direction="horizontal" alignment="center" :size="200">
            <el-button type="primary" size="medium" autocomplete="off" @click="register"
              >注册</el-button
            >
            <el-button
              type="warning"
              size="medium"
              autocomplete="off"
              @click="$router.push('/login')"
              >返回登录</el-button
            >
          </el-space>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'user_register',
  data() {
    return {
      user: {},
      rules: {
        email: [
            { required: true, message: '邮箱不能为空', trigger: 'blur' },
            { type: 'email', message: '请输入正确的邮箱地址', trigger: 'blur' },
        ],
        phone: [
          { required: true, message: '手机号不能为空', trigger: 'blur' },
          { min: 11, max: 11, message: '长度应为 11 个数字', trigger: 'blur' },
          { pattern: /^1[34578]\d{9}$/, message: "请输入正确的手机号（目前只支持中国大陆）", trigger: "blur",},
        ],
        id_number: [
          { required: true, message: '身份证号不能为空', trigger: 'blur' },
          { min: 18, max: 18, message: '长度应为 18 个字符', trigger: 'blur' },
          { pattern:/(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)/, message: '身份证格式不正确', trigger: 'blur' }
        ],
        username: [
          { required: true, message: '用户名不能为空', trigger: 'blur' },
          { min: 3, max: 10, message: '长度在 3 到 10 个字符', trigger: 'blur' },
          { pattern: /^(\w){3,10}$/, message: '用户名只能包含字母、数字或下划线'}
        ],
        password: [
          { required: true, message: '密码不能为空', trigger: 'blur' },
          { min: 6, max: 32, message: '长度在 6 到 32 个字符', trigger: 'blur' },
          { pattern: /^((?=.*\d)(?=.*[a-zA-Z]))|((?=.*[-_])(?=.*\d))|((?=.*[a-zA-Z])(?=.*[-_])).{6,32}$/,
            message: '字⺟，数字或者特殊字符（-_）⾄少包含两种'}
        ],
        user_type: [{ required: true, message: '请选择要注册的用户类型', trigger: 'change' }],
        confirmPassword: [
          { required: true, message: '密码不能为空', trigger: 'blur' },
          { min: 6, max: 32, message: '长度在 6 到 32 个字符', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    register() {
      this.$refs['userForm'].validate((valid) => {
        if (valid) {
          // 表单校验合法
          if (this.user.password !== this.user.confirmPassword) {
            this.$message.error('两次输入的密码不一致')
            return false
          }
          this.request.post('/user-data/register', this.user).then((res) => {
            if (res.code === '200') {
              localStorage.setItem('user_data', JSON.stringify(res.data)) // 存储用户信息到浏览器
              if (res.data.userType === 'user') {
                this.$router.push('/user')
              } else if (res.data.userType === 'seller') {
                this.$router.push('/seller')
              } else {
                this.$router.push('/')
              }
              this.$message.success('注册成功，自动登录')
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

<style>
.wrapper {
  height: 100vh;
  background-color: #72767b;
  overflow: hidden;
}
</style>
