<template>
  <div>
    <el-select v-model="search_info.all" placeholder="是否指定用户种类" style="width: 180px;">
      <el-option label="默认所有用户" value= 'true' />
      <el-option label="指定用户种类" value= 'false' />
    </el-select>
    <el-select v-model="search_info.gender" placeholder="性别"  style="width: 120px">
      <el-option label="不区分" value= 'null' />
      <el-option label="男" value= 'true' />
      <el-option label="女" value = 'false' />
    </el-select>
    <el-input-number
      v-model="search_info.age_start"
      style="width: 200px"
      :min="0"
      :max="150"
      controls-position="right"
      placeholder="请输入起始年龄"
    />
    <el-input-number
      v-model="search_info.age_end"
      style="width: 200px"
      :min="0"
      :max="150"
      controls-position="right"
      placeholder="请输入结束年龄"
    />
    <el-button @click="renderChart">
      <el-icon><Search /></el-icon>
    </el-button>
  </div>
  <div style="margin-top: 20px">
    <canvas id="myChart" width="150" height="60"></canvas>
  </div>
</template>

<script>
import { Chart } from 'chart.js/auto';
import { Search } from "@element-plus/icons-vue";

export default {
  name: "statistics_favorite",
  components: { Search },
  data() {
    return {
      chartData: {
        labels: [],
        datasets: [
          {
            label: '',
            backgroundColor: [
              'rgba(255, 99, 132, 0.2)',
              'rgba(255, 159, 64, 0.2)',
              'rgba(255, 205, 86, 0.2)',
              'rgba(75, 192, 192, 0.2)',
              'rgba(54, 162, 235, 0.2)',
              'rgba(153, 102, 255, 0.2)',
              'rgba(201, 203, 207, 0.2)'
            ],
            borderColor: [
              'rgb(255, 99, 132)',
              'rgb(255, 159, 64)',
              'rgb(255, 205, 86)',
              'rgb(75, 192, 192)',
              'rgb(54, 162, 235)',
              'rgb(153, 102, 255)',
              'rgb(201, 203, 207)'
            ],
            borderWidth: 1,
            data: [],
          },
        ],
      },
      search_info:{
        all:null,
        gender: null,
        age_start: null,
        age_end: null,
      },
    };
  },
  created() {
    this.search_info.all='true'
  },
  mounted() {
    this.renderChart();
  },
  methods: {
    renderChart() {
      const ctx = document.getElementById('myChart').getContext('2d');
      // 销毁之前的图表实例
      if (this.chartInstance) {
        this.chartInstance.destroy();
      }

      if(this.search_info.all!=null){
        this.SearchStatisticsFavorite()
      }

      // 创建新的图表实例
      this.chartInstance = new Chart(ctx, {
        type: 'bar',
        data: this.chartData,
        options: {
          scales: {
            y: {
              beginAtZero: true,
            },
          },
        },
      });
    },
    SearchStatisticsFavorite(){
      this.search_info.all = this.search_info.all === 'true'
      if(this.search_info.all === true){
        this.search_info.gender=null
        this.search_info.age_start=null
        this.search_info.age_end=null
        // request
        this.request.post("/favorite/statistics",this.search_info).then((res) => {
          if (res.status === 200) {
            if(res.data.length<1){
              this.$message.warning('数据为空');
            }else {
              this.chartData.labels=[];
              this.chartData.datasets=[
                {
                  label: '所有用户收藏前10的商品',
                  backgroundColor: [
                    'rgba(255, 99, 132, 0.2)',
                    'rgba(255, 159, 64, 0.2)',
                    'rgba(255, 205, 86, 0.2)',
                    'rgba(75, 192, 192, 0.2)',
                    'rgba(54, 162, 235, 0.2)',
                    'rgba(153, 102, 255, 0.2)',
                    'rgba(201, 203, 207, 0.2)'
                  ],
                  borderColor: [
                    'rgb(255, 99, 132)',
                    'rgb(255, 159, 64)',
                    'rgb(255, 205, 86)',
                    'rgb(75, 192, 192)',
                    'rgb(54, 162, 235)',
                    'rgb(153, 102, 255)',
                    'rgb(201, 203, 207)'
                  ],
                  borderWidth: 1,
                  data: [],
                },
              ];
              res.data.forEach(item => {
                this.chartData.labels.push(item.CommodityItem.item_name);
              });
              res.data.forEach(item => {
                this.chartData.datasets[0].data.push(item.count);
              });
            }

          }
        }).catch(error => {
          this.$message.error(error);
        });
      }else {
        if(this.search_info.gender==null||this.search_info.age_start==null||this.search_info.age_end==null){
          this.$message.error('若指定用户范围，请指定完整的性别和年龄范围')
        }else{
          if(this.search_info.gender === 'null'){
            this.search_info.gender=null
          }else {
            this.search_info.gender = this.search_info.gender === 'true'
          }
          this.search_info.age_start=parseInt(this.search_info.age_start)
          this.search_info.age_end=parseInt(this.search_info.age_end)
          let label = '性别:'+ (this.search_info.gender==null?'不区分':this.search_info.gender===true?'男':'女') + ' 年龄:'+this.search_info.age_start+'~'+this.search_info.age_end+' 用户收藏前10的商品'
            // request
            this.request.post("/favorite/statistics", this.search_info).then((res) => {
              if (res.status === 200) {
                if(res.data.length<1){
                  this.$message.warning('数据为空');
                }else {
                  this.chartData.labels = [];
                  this.chartData.datasets = [
                    {
                      label: label,
                      backgroundColor: [
                        'rgba(255, 99, 132, 0.2)',
                        'rgba(255, 159, 64, 0.2)',
                        'rgba(255, 205, 86, 0.2)',
                        'rgba(75, 192, 192, 0.2)',
                        'rgba(54, 162, 235, 0.2)',
                        'rgba(153, 102, 255, 0.2)',
                        'rgba(201, 203, 207, 0.2)'
                      ],
                      borderColor: [
                        'rgb(255, 99, 132)',
                        'rgb(255, 159, 64)',
                        'rgb(255, 205, 86)',
                        'rgb(75, 192, 192)',
                        'rgb(54, 162, 235)',
                        'rgb(153, 102, 255)',
                        'rgb(201, 203, 207)'
                      ],
                      borderWidth: 1,
                      data: []
                    }
                  ];
                  res.data.forEach(item => {
                    this.chartData.labels.push(item.CommodityItem.item_name);
                  });
                  res.data.forEach(item => {
                    this.chartData.datasets[0].data.push(item.count);
                  });
                }
              }
            }).catch(error => {
              this.$message.error(error.response.data.message);
            });
        }
      }
      this.search_info.all = null
      this.search_info.gender=null
      this.search_info.age_start=null
      this.search_info.age_end=null
    },
  },
};
</script>
