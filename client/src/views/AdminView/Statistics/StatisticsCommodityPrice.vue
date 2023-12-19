<template>
  <div>
    <el-input-number
      v-model="commodity_id"
      style="width: 250px"
      :min="0"
      controls-position="right"
      placeholder="请输入可发布商品的序号"
    />
    <el-date-picker
      style="margin-left: 30px"
      v-model="range"
      type="daterange"
      unlink-panels
      range-separator="到"
      start-placeholder="起始时间"
      end-placeholder="结束时间"
      value-format="YYYY-MM-DD hh:mm:ss"
      :shortcuts="shortcuts"
    />
    <el-button style="margin-left: 30px" @click="renderChart">
      <el-icon><Search /></el-icon>
    </el-button>
    <div style="float: right">
      <span style="color: #0059ff">价格最高：{{this.chartData.labels[result0.maxIndex]}}</span>
      <span style="color: rgba(59,89,134,0.79); margin-left: 40px;">价格最低：{{this.chartData.labels[result0.minIndex]}}</span>
    </div>
    <div style="float: right">
      <span style="color: #00f504">价格差最高：{{this.chartData.labels[result1.maxIndex]}}</span>
      <span style="color: rgba(103,194,58,0.53); margin-left: 10px;">价格差最低：{{this.chartData.labels[result1.minIndex]}}</span>
    </div>
  </div>
  <div style="margin-top: 20px">
    <canvas id="myChart" width="150" height="60"></canvas>
  </div>
</template>

<script>
import { Chart } from 'chart.js/auto';
import { Search } from "@element-plus/icons-vue";

export default {
  name: "statistics_price_seller",
  components: { Search },
  data() {
    return {
      chartData: {
        labels: [],
        datasets: [
          {
            label: '',
            backgroundColor: [
              'rgba(54, 162, 235, 0.2)',
            ],
            borderColor: [
              'rgb(54, 162, 235)',
            ],
            borderWidth: 1,
            data: [],
          },
          {
            label: '',
            backgroundColor: [
              'rgba(33,245,0,0.32)',
            ],
            borderColor: [
              'rgba(33,245,0,0.7)',
            ],
            borderWidth: 1,
            data: [],
          },
        ],
      },
      commodity_id: null,
      range:null,
      search:{
        commodity_id:null,
        time_start:null,
        time_end: null,
      },
      shortcuts: [
        {
          text: "近一周",
          value: function() {
            const end = new Date();
            const start = new Date();
            start.setDate(end.getDate() - 7);
            return [start, end];
          }
        },
        {
          text: "近一月",
          value: function() {
            const end = new Date();
            const start = new Date();
            start.setMonth(end.getMonth() - 1);
            return [start, end];
          }
        },
        {
          text: "近一年",
          value: function() {
            const end = new Date();
            const start = new Date();
            start.setFullYear(end.getFullYear() - 1);
            return [start, end];
          }
        }
      ],
    };
  },
  mounted() {
    this.renderChart();
  },
  computed: {
    result0() {
      return this.calculateMinMaxIndex(this.chartData.datasets[0].data);
    },
    result1() {
      return this.calculateMinMaxIndex(this.chartData.datasets[1].data);
    },
  },
  methods: {
    renderChart() {
      const ctx = document.getElementById('myChart').getContext('2d');
      // 销毁之前的图表实例
      if (this.chartInstance) {
        this.chartInstance.destroy();
      }

      if(this.commodity_id!=null){
        this.SearchStatisticsPriceSeller()
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
    SearchStatisticsPriceSeller(){
      if(this.commodity_id==null || this.range==null){
        this.$message.error('请输入商品ID并指定时间范围');
      }else {
        this.search.commodity_id=parseInt(this.commodity_id)
        this.search.time_start=this.range[0]
        this.search.time_end=this.range[1]
        this.request.post("/price/statistics",this.search).then((res) => {
          if (res.status === 200) {
            if(res.data.length<1){
              this.$message.warning('数据为空');
            }else {
              this.chartData.labels=[];
              this.chartData.datasets=[
                {
                  label: res.data[0].CommodityItem.Commodity.default_name+" 在不同商家的价格",
                  backgroundColor: [
                    'rgba(54, 162, 235, 0.2)',
                  ],
                  borderColor: [
                    'rgb(54, 162, 235)',
                  ],
                  borderWidth: 1,
                  data: [],
                },
                {
                  label: res.data[0].CommodityItem.Commodity.default_name+" 在不同商家的价格差",
                  backgroundColor: [
                    'rgba(33,245,0,0.32)',
                  ],
                  borderColor: [
                    'rgba(33,245,0,0.7)',
                  ],
                  borderWidth: 1,
                  data: [],
                },
              ];
              res.data.forEach(item => {
                this.chartData.labels.push(item.CommodityItem.Seller.username);
              });
              res.data.forEach(item => {
                this.chartData.datasets[0].data.push(item.CommodityItem.price);
              });
              res.data.forEach(item => {
                this.chartData.datasets[1].data.push(item.price_variance);
              });
            }

          }
        });
      }
    },
    calculateMinMaxIndex(arr) {
      if (arr.length === 0) {
        return { maxIndex: null, minIndex: null };
      }

      let maxIndex = 0;
      let minIndex = 0;
      let maxValue = arr[0];
      let minValue = arr[0];

      for (let i = 1; i < arr.length; i++) {
        if (arr[i] > maxValue) {
          maxValue = arr[i];
          maxIndex = i;
        }

        if (arr[i] < minValue) {
          minValue = arr[i];
          minIndex = i;
        }
      }

      return {
        maxIndex: maxIndex,
        minIndex: minIndex,
      };
    },
  },
};
</script>
