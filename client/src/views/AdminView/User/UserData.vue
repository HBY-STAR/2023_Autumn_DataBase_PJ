<template>
  <el-scrollbar height="605px">
    <el-table :data="all_user_data" height="605px" border="true">
      <el-table-column label="用户序列号" prop="uid" width="120"/>
      <el-table-column label="用户名" prop="username" width="120"/>
      <el-table-column label="密码" prop="password" width="140"/>
      <el-table-column label="邮箱" prop="email" width="240"/>
      <el-table-column label="电话" prop="phone" width="120"/>
      <el-table-column label="用户类型" prop="userType" width="120"/>
      <el-table-column label="身份证号" prop="idNumber" width="200"/>
    </el-table>
  </el-scrollbar>
</template>

<script>
export default {
  name:"all_user_data",
  data(){
    return{
      all_user_data: localStorage.getItem('all_user_data')
        ? JSON.parse(localStorage.getItem('all_user_data'))
        : {}
    }
  },
  created() {
    // 从后台获取最新的User数据
    this.findAll()
  },
  methods:{
    findAll(){
      this.request.get('/user-data/find_all_user_data').then((res) => {
        if (res.code === '200') {
          localStorage.setItem('all_user_data', JSON.stringify(res.data)) // 存储用户信息到浏览器
        } else {
          this.$message.error(res.msg)
        }
      })
    }
  }
}
</script>