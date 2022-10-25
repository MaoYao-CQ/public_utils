<template>
  <div>
    <div>
      <div class="border-bottom">
        <form-search
            :data-search="dataSearch"
            :alert-search-data="alertSearchData"
        >
          <template v-slot:append>
            <el-button type="info" size="mini" @click="dialogVirusDataImportOpen">导入</el-button>
            <el-button type="info" size="mini" @click="virusDataExport">导出</el-button>
            <el-button type="info" size="mini" @click="virusDataFiles">查看文件</el-button>
          </template>
        </form-search>
      </div>
      <div>
        <el-table :data="entities"
                  style="width: 100%"

        >
          <el-table-column type="index" width="40"></el-table-column>
          <!---->
          <el-table-column label="任务">
            <template v-slot="{row}">
              {{ $LeUtils.valueGet(row, 'task.name') }}
            </template>
          </el-table-column>

          <el-table-column label="病毒">
            <template v-slot="{row}">
              [{{ $LeUtils.valueGet(row, 'virus.name') }}] {{ $LeUtils.valueGet(row, 'virus.id') }}
            </template>
          </el-table-column>

          <el-table-column label="事件名">
            <template v-slot="{row}">
              {{ $LeUtils.valueGet(row, 'task_name') }}
            </template>
          </el-table-column>
          <el-table-column label="事件状态">
            <template v-slot="{row}">
              {{ $LeUtils.valueGet(row, 'task_status') }}
            </template>
          </el-table-column>
          <el-table-column label="事件结果">
            <template v-slot="{row}">
              {{ $LeUtils.valueGet(row, 'task_result') }}
            </template>
          </el-table-column>
          <el-table-column label="上报IP">
            <template v-slot="{row}">
              {{ $LeUtils.valueGet(row, 'client_ip') }}
            </template>
          </el-table-column>
          <el-table-column label="上报时间">
            <template v-slot="{row}">
              {{ $LeUtils.formatDataTime($LeUtils.valueGet(row, 'created_at')) }}
            </template>
          </el-table-column>

          <el-table-column label="平台同步" width="80">
            <template v-slot="{row}">
              <el-tag v-if="row.synced" type="success" size="small">已同步</el-tag>
              <el-tag v-else type="info" size="small">未同步</el-tag>
            </template>
          </el-table-column>

        </el-table>
      </div>
      <div class="mt-2">
        <div class="float-end">
          <pagination :data-pager="dataPager" page-layout="total, prev, pager, next">

          </pagination>
        </div>
      </div>
    </div>

    <el-dialog
        title="导出文件列表"
        :visible.sync="dialogVirusDataExportFilesVisible"
        width="75%">
      <div>
        <el-table :data="virusDataExportFiles ||[]"
                  style="width: 100%"

        >
          <el-table-column label="文件名">
            <template v-slot="{row}">
              <el-link :href="'/console/virus-data/exported-file-download/'+ row.name">{{row.name}}</el-link>
            </template>
          </el-table-column>
          <el-table-column label="导出时间" prop="mod_time"></el-table-column>
          <el-table-column label="文件大小" prop="size"></el-table-column>
        </el-table>
      </div>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogVirusDataExportFilesCLose">取 消</el-button>
      </div>
    </el-dialog>

    <el-dialog
        title="导入数据"
        :visible.sync="dialogVirusDataImportVisible"
        width="75%">
      <div>
        <el-form label-width="100px">
          <el-form-item label="选择文件">
            <el-select v-model="virusDataImportFile" placeholder="请选择">
              <el-option
                  v-for="item in (virusDataExportFiles ||[])"
                  :key="item.name"
                  :label="item.name+'['+ item.mod_time + ',' + item.size +']'"
                  :value="item.name">
              </el-option>
            </el-select>
          </el-form-item>
        </el-form>
      </div>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogVirusDataImportCLose">取 消</el-button>
        <el-button @click="virusDataImport">导 入</el-button>
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
      type: Object,
      default: function () {
        return {};
      }
    }
  },
  data() {
    return {
      dialogEditVisible: false,
      currentEditEntity: false,
      entities: {},
      pager: {},
      searchData: {},
      dialogVirusDataImportVisible: false,
      virusDataImportFile: "",
      dialogVirusDataExportFilesVisible: false,
      virusDataExportFiles: [],
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
    alertSearchData:function (data){
      this.searchData = data
      return data
    },
    virusDataExport: function () {
      let _this = this;
      LeRoute.request({
        uri: "/console/virus-data/export",
        method: "post"
      }, {
        task_id : LeUtils.valueGet(this.searchData,"task_id"),
        virus_id : LeUtils.valueGet(this.searchData,"virus_id"),
      }).then(function (res_data) {
        _this.$message.success({
          message: LeUtils.valueGet(res_data, "message", "执行成功"),
          duration: 5000
        })
      })
    },
    virusDataFiles: function () {
      let _this = this;
      LeRoute.request({
        uri: "/console/virus-data/exported-files",
        method: "get"
      }, {}).then(function (res_data) {
        _this.dialogVirusDataExportFilesOpen(LeUtils.valueGet(res_data,"data",[]))
      })
    },
    dialogVirusDataExportFilesOpen: function (files) {
      this.dialogVirusDataExportFilesVisible = false
      this.virusDataExportFiles = files
      this.dialogVirusDataExportFilesVisible = true
    },
    dialogVirusDataExportFilesCLose: function () {
      this.dialogVirusDataExportFilesVisible = false
      this.virusDataExportFiles = []
    },
    dialogVirusDataImportOpen: function (files) {
      this.dialogVirusDataImportVisible = false
      let _this = this
      LeRoute.request({
        uri: "/console/virus-data/exported-files",
        method: "get"
      }, {}).then(function (res_data) {
        _this.virusDataExportFiles = LeUtils.valueGet(res_data,"data",[])
        _this.dialogVirusDataImportVisible = true
      })

    },
    dialogVirusDataImportCLose: function () {
      this.dialogVirusDataImportVisible = false
      this.virusDataExportFiles = []
    },

    virusDataImport: function () {
      let _this = this;
      LeRoute.request({
        uri: "/console/virus-data/import",
        method: "post"
      }, {file_name:this.virusDataImportFile}).then(function (res_data) {
        _this.dialogVirusDataImportCLose()
        _this.$message.success({
          message: LeUtils.valueGet(res_data, "message", "执行成功"),
          duration: 5000
        })


      })
    },
  }
};
</script>
