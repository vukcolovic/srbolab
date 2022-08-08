<template>
  <div class="container">
    <div class="row m-3">
      <div class="btn-group">
        <button class="iconBtn" title="Dodaj">
          <i class="fa fa-plus"></i>
        </button>
        <button class="iconBtn" title="Pregledaj" :disabled="selectedIrregularity == null" >
          <i class="fa fa-eye"></i>
        </button>
        <button class="iconBtn" title="Izmeni" :disabled="selectedIrregularity == null" >
          <i class="fa fa-pencil">
          </i>
        </button>
      </div>
    </div>
    <div class="row m-3">
      <vue-table-lite
          @row-clicked="selectedIrregularity"
          :columns="columns"
          :rows="rows"
          :total= "totalCount"
          @do-search="getAll"
          :is-loading="isLoading"
      ></vue-table-lite>
    </div>
  </div>
</template>

<script>
import VueTableLite from "vue3-table-lite";
import axios from "axios";
import notie from 'notie';

export default {
  name: 'IrregularityLevels',
  components: { VueTableLite },
  data() {
    return {
      columns: [
        {
          label: 'ID',
          field: 'id',
          width: '3%',
          isKey: true,
        },
        {
          label: 'Naziv',
          field: 'code',
          width: '10%',
        },
        {
          label: 'Kreiran',
          field: 'created_at',
          width: '10%',
        }
      ],
      rows: [],
      selectedIrregularity: null,
      isLoading: false,
      totalCount: 0
    }
  },
  methods: {
    selectLevel(rowData) {
      this.selectedIrregularity = rowData;
    },
    async getAll() {
      this.isLoading = true;
      await axios.get('/enumeration/irregularity-levels/all').then((response) => {
        if (response.data === null || response.data.Status === 'error') {
          notie.alert({
            type: 'error',
            text: 'Greska: ' + response.data.ErrorMessage,
            position: 'bottom',
          })
          return;
        }
        this.rows = JSON.parse(response.data.Data);
      }, (error) => {
        notie.alert({
          type: 'error',
          text: "Greska: " + error,
          position: 'bottom',
        })
      });

      this.isLoading = false;
    },
  },
  async created() {
    await this.getAll();
    this.totalCount = this.rows.length;
  }
}
</script>

<style scoped>
::v-deep(.vtl-table .vtl-thead .vtl-thead-th) {
  font-size: 12px;
}
::v-deep(.vtl-table td),
::v-deep(.vtl-table tr) {
  font-size: 12px;
  padding: 5px;
}
::v-deep(.vtl-paging-info) {
  font-size: 12px;
  padding: 5px;
}
::v-deep(.vtl-paging-count-label),
::v-deep(.vtl-paging-page-label),
::v-deep(.vtl-paging-count-dropdown),
::v-deep(.vtl-paging-page-dropdown){
  font-size: 12px;
  padding: 5px;
}
::v-deep(.vtl-paging-pagination-page-link) {
  font-size: 12px;
  padding: 5px;
}
</style>
