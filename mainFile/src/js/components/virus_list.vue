<template>
  <div>
    <div>
      <el-descriptions class="margin-top" title="" :column="1"  border>
        <el-descriptions-item label="概况">
          <div v-html="virusSummary(entities)"></div>
        </el-descriptions-item>
      </el-descriptions>
      <el-table :data="entities"
                style="width: 100%"

      >
        <el-table-column type="index" width="40"></el-table-column>
        <!---->
        <el-table-column label="病毒文件" width="350">
          <template v-slot="{row}">
            {{row.name}}: {{ row.id}} <br>
            HASH: {{ row.file_hash}}
          </template>
        </el-table-column>

        <el-table-column label="运行系统" width="90">
          <template v-slot="{row}">
            {{ row.run_os }}
          </template>
        </el-table-column>

        <el-table-column label="大小" width="90">
          <template v-slot="{row}">
            {{$LeUtils.formatFileSize(row.file_size)}}
          </template>
        </el-table-column>

        <el-table-column label="状态">
          <template v-slot="{row}">
            <div v-if="($LeUtils.valueGet(row,'task',null) || ('used' == $LeUtils.valueGet(row,'status','')))" >
              <el-tag  size="small" type="success">
                已分配
              </el-tag>
              <span class="text-nowrap">[{{$LeUtils.valueGet(row,'task.name','')}}] </span>
              <span class="text-nowrap">[{{$LeUtils.valueGet(row,'task.executor','')}}]</span>
            </div>

            <el-tag v-else  size="small" type="warning">
              未分配
            </el-tag>
          </template>
        </el-table-column>


        <el-table-column label="操作" align="right" width="">
          <template v-slot="{row}">
            <el-link
                type="primary"
                size="mini"
                class="p-1"
                :href="'/console/virus/entity/'+row.id"
            >编辑
            </el-link>

            <el-link
                type="info"
                size="mini"
                class="p-1"
                :href="'/console/virus-data/index?virus_id='+row.id"
            >查看日志
            </el-link>
          </template>
        </el-table-column>
      </el-table>
    </div>


  </div>
</template>

<script>

export default {
  components: {},
  props: {
    dataEntities: {
      type: Array,
      default: function () {
        return [];
      }
    },
    virtualizationOptions: {
      type: Array,
      default: function () {
        return [];
      }
    },

    dataSearch: {
      type: Object,
      default: function () {
        return {};
      }
    },

    apiRoutes: {
      type: Object,
      default: function () {
        return {};
      }
    },
    dataPager: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      dialogEditVisible: false,
      currentEditEntity: false,
      entities: {},
      pager: {},
    };
  },
  async created() {
    let _this = this;
    _this.entities = _this.dataEntities;
    _this.pager = _this.dataPager;
  },
  mounted() {
    let _this = this;
  },

  updated() {
    let _this = this;


  },

  methods: {

    typeIs: function (type, obj) {
      return LeUtils.typeIs(type, obj);
    },

    deleteItem: function (item) {
      let _this = this;

      let delete_api = LeRoute.get(_this.apiRoutes.delete || '');
      let entity_id = LeUtils.valueGet(item, 'id', '');

      if (!delete_api || !entity_id) {
        _this.$message.error('数据有误');
      }

      this.$confirm('请确认要删除授权？', '警告！', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        LeRoute.request(delete_api, {id: entity_id}).then(function (res_data) {
          _this.$message.success('删除成功');
          window.location.reload();
        }).catch(function (error) {
          _this.$message.error('删除失败');
        });
      });
    },

    virusSummary: function (virus) {
      let total = 0
      let notImplanted = 0
      let implanted = 0
      let activated = 0
      let finished = 0
      let deleted = 0
      if (LeUtils.typeIs('array', virus)) {
        for (let i in virus) {
          let vs = LeUtils.valueGet(virus[i], 'virus_status')
          total += 1
          if ("deleted" == vs) {
            deleted += 1
            finished += 1
            activated += 1
            implanted += 1
          } else if ("finished" == vs) {
            finished += 1
            activated += 1
            implanted += 1
          } else if ("execute_command" == vs) {
            activated += 1
            implanted += 1
          } else if ("run_program" == vs) {
            implanted += 1
          } else {
            notImplanted += 1
          }
        }

      }
      let s = ''
      s += "<span class='text-nowrap'>总数:<strong>" + total + "</strong>,</span>"
      s += "<span class='text-nowrap'>待植:<strong>" + notImplanted + "</strong>,</span>"
      s += "<span class='text-nowrap'>已植:<strong>" + implanted + "</strong>,</span>"
      s += "<span class='text-nowrap'>激活:<strong>" + activated + "</strong>,</span>"
      s += "<span class='text-nowrap'>完成:<strong>" + finished + "</strong>,</span>"
      s += "<span class='text-nowrap'>自毁:<strong>" + deleted + "</strong></span>"
      return s
    },
  }
};
</script>
