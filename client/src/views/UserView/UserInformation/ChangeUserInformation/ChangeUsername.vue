<template>
  <div style="margin-top: 10px">
    <el-form :model="user" :rules="rules" ref="userForm">
    <el-form-item prop="username">
      <el-input
          v-model="user.username"
          placeholder="请输入用户名"
      ></el-input>
    </el-form-item>
    </el-form>
    <div style="margin-top: 30px">
      <el-button autocomplete="off" type="primary" @click="changeUsername"
      >保存
      </el-button
      >
      <el-button
          autocomplete="off"
          style="position: absolute;right: 0;"
          type="warning"
          @click="$router.push('/personal_information')"
      >取消
      </el-button
      >
    </div>
  </div>
</template>

<script>

export default {
  name: "change_username",
  data() {
    return {
      user:{
        username: '',
        uid: 0,
      },
      rules:{
        username: [
          { required: true, message: '用户名不能为空', trigger: 'blur' },
          { min: 3, max: 10, message: '长度在 3 到 10 个字符', trigger: 'blur' },
          { pattern: /^(\w){3,10}$/, message: '用户名只能包含字母、数字或下划线', trigger: 'blur'}
        ],
      },
      user_data: localStorage.getItem('user_data')
          ? JSON.parse(localStorage.getItem('user_data'))
          : {},
    }
  },
  methods: {
    changeUsername() {
      this.$refs['userForm'].validate((valid) => {
        if (valid) {
          this.user.uid = this.user_data.uid;
          this.request.post('/user-data/change_username', this.user).then((res) => {
            if (res.code === '200') {
              localStorage.setItem('user_data', JSON.stringify(res.data)) // 存储用户信息到浏览器
              this.$message.success("保存成功")
            } else {
              this.$message.error(res.msg)
            }
          })
        }
      })
    },
  }
}
</script>