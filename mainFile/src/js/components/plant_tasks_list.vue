<template>
  <div>
    <div>
      <el-table :data="entities"
                style="width: 100%"
      >
        <el-table-column type="index" width="40"></el-table-column>

        <el-table-column label="任务编号" prop="code">
        </el-table-column>
        <el-table-column label="任务名称" prop="name">
        </el-table-column>
        <el-table-column label="执行单位" prop="organization">
        </el-table-column>
        <el-table-column label="执行计划" width="380">
          <template v-slot="{row}">
            <el-descriptions class="margin-top" :column="2" size="mini">
              <el-descriptions-item label="植入">
                {{ row.implant_time }}
              </el-descriptions-item>
              <el-descriptions-item label="激活">
                {{ row.activated_time }}
              </el-descriptions-item>
              <el-descriptions-item label="结束">
                {{ row.end_time }}
              </el-descriptions-item>
              <el-descriptions-item label="销毁">
                {{ row.self_delete_time }}
              </el-descriptions-item>

            </el-descriptions>
          </template>
        </el-table-column>
        <el-table-column label="回连地址" prop="service_ip">
        </el-table-column>


        <el-table-column label="操作">
          <template v-slot="{row}">
            <el-link
                type="primary"
                size="mini"
                class="p-1"
                :href="'/console/plant-tasks/entity/'+row.id"
            >修改配置
            </el-link>
            <el-link
                type="info"
                size="mini"
                class="p-1"
                :href="'/console/virus-data/index?task_id='+row.id"
            >查看日志
            </el-link>

          </template>
        </el-table-column>

        <el-table-column label="病毒概况" width="180">
          <template v-slot="{row}">
            <span class="fw-bolder pe-2" v-html="taskVirusSummary(row.virus)">
            </span>

          </template>
        </el-table-column>
        <el-table-column type="expand">

          <template v-slot="{row}">
            <div class="ms-5">
              <el-table :data="row.virus || []"
                        size="mini"
                        border

              >
                <el-table-column type="index" width="40"></el-table-column>
                <el-table-column label="病毒" width="300">
                  <template v-slot="{row}">
                    [{{ row.name }}] {{ row.id }}
                  </template>
                </el-table-column>
                <el-table-column label="执行人" prop="executor" width="80">
                </el-table-column>
                <el-table-column label="网络区域" prop="network_area" width="100">
                </el-table-column>
                <el-table-column label="目标IP" prop="target_ip">
                </el-table-column>
                <el-table-column label="目标MAC" prop="target_mac">
                </el-table-column>
                <el-table-column label="日志路径" prop="log_path">
                </el-table-column>
                <el-table-column label="状态" prop="run_status" width="80">
                  <template v-slot="{row}">
                    <el-tag type="warning" size="mini" v-if="'deleted'==row.virus_status">
                      已自毁
                    </el-tag>
                    <el-tag type="warning" size="mini" v-else-if="'finished'==row.virus_status">
                      已完成
                    </el-tag>
                    <el-tag type="success" size="mini" v-else-if="'execute_command'==row.virus_status">
                      已激活
                    </el-tag>
                    <el-tag type="info" size="mini" v-else-if="'run_program'==row.virus_status">
                      已植入
                    </el-tag>
                    <el-tag type="danger" size="mini" v-else>
                      待植入
                    </el-tag>
                  </template>
                </el-table-column>
                <el-table-column label="最后回连时间" prop="last_log_time">
                </el-table-column>
                <el-table-column label="操作" width="220">
                  <template v-slot="socpVirus">
                    <el-button
                        type="primary"
                        size="mini"
                        class="p-1"
                        @click="openDialogAddVirus(row,socpVirus.row)"
                    >修改
                    </el-button>
                    <el-button
                        type="info"
                        size="mini"
                        class="p-1"
                        @click="downloadVirus(row,socpVirus.row)"
                    >下载
                    </el-button>
                    <el-link
                        type="info"
                        size="mini"
                        class="p-1"
                        :href="'/console/virus-data/index?virus_id='+socpVirus.row.id"
                    >查看日志
                    </el-link>
                  </template>

                </el-table-column>
              </el-table>
              <div class="mt-2">
                <el-button
                    type="primary"
                    size="mini"
                    class="p-1 "
                    plain
                    @click="openDialogAddVirus(row)"
                >分配病毒
                </el-button>
              </div>
            </div>
          </template>
        </el-table-column>

      </el-table>
    </div>

    <el-dialog
        title="添加病毒"
        v-if="dialogAddVirusVisible"
        :visible.sync="dialogAddVirusVisible"
        show-close
        :destroy-on-close="true"
        :before-close="closeDialogAddVirus"
        width="75%">
      <div>
        <form-plant-tasks-virus
            :data-task="currentEditEntity"
            :data-entity="currentVirusEntity"
        >
        </form-plant-tasks-virus>
      </div>
    </el-dialog>
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

    dataSearch: {
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
  computed: {},
  data() {
    return {
      dialogAddVirusVisible: false,
      currentEditEntity: null,
      currentVirusEntity: null,
      entities: {},
      pager: {},
    };
  },
  async created() {
    let _this = this;
    _this.entities = _this.dataEntities;
    _this.pager = _this.dataPager;
    //console.log(this.entities)
  },
  mounted() {
    let _this = this;
    //console.log(this.dataEntities)
  },

  updated() {
    let _this = this;


  },

  methods: {

    typeIs: function (type, obj) {
      return LeUtils.typeIs(type, obj);
    },

    taskVirusSummary: function (virus) {
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

    downloadVirus: function (task, virus) {
      let _this = this
      LeRoute.request({
        uri: "/console/plant-tasks/down-virus/{id}",
        method: "get"
      }, {id: task.id, virus_id: virus.id}).then(function (res_data) {
        if (!LeUtils.valueGet(res_data, 'status')) {
          return
        }
        let text = '';
        text += '<label class="me-2">下载地址:</label>';
        text += '<a target="_blank" href="' + LeUtils.valueGet(res_data, 'data.url', "") + '">点击下载</a><br>';
        text += '<label class="me-2">运行参数:</label><br><textarea class="w-100" rows="10" readonly="readonly">';
        text += LeUtils.valueGet(res_data, 'data.params', "");
        text += '</textarea>';


        _this.$alert(text, '病毒下载', {
          confirmButtonText: '关闭',
          dangerouslyUseHTMLString: true,
          customClass: 'w-50',
        });


      })
    },
    openDialogAddVirus: function (task, virus) {
      this.currentEditEntity = task
      this.currentVirusEntity = virus
      this.dialogAddVirusVisible = true
    },
    closeDialogAddVirus: function () {
      this.currentEditEntity = null
      this.currentVirusEntity = null
      this.dialogAddVirusVisible = false
    },
  }
};
</script>
