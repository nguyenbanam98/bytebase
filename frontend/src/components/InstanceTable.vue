<template>
  <BBTable
    :columnList="state.columnList"
    :dataSource="instanceList"
    :showHeader="true"
    :leftBordered="false"
    :rightBordered="false"
    @click-row="clickInstance"
  >
    <template v-slot:body="{ rowData: instance }">
      <BBTableCell :leftPadding="4" class="w-8">
        <img
          v-if="instance.engine == 'MYSQL'"
          class="h-10 w-auto"
          src="../assets/db-mysql.svg"
        />
      </BBTableCell>
      <BBTableCell class="w-32">
        {{ instanceName(instance) }}
      </BBTableCell>
      <BBTableCell class="w-16">
        {{ environmentNameFromId(instance.environment.id) }}
      </BBTableCell>
      <BBTableCell class="w-48">
        <template v-if="instance.port"
          >{{ instance.host }}:{{ instance.port }}</template
        ><template v-else>{{ instance.host }}</template>
      </BBTableCell>
      <BBTableCell class="w-4">
        <button
          v-if="instance.externalLink?.trim().length != 0"
          class="btn-icon"
          @click.stop="window.open(urlfy(instance.externalLink), '_blank')"
        >
          <svg
            class="w-4 h-4"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"
            ></path>
          </svg>
        </button>
      </BBTableCell>
      <BBTableCell class="w-16">
        {{ humanizeTs(instance.createdTs) }}
      </BBTableCell>
    </template>
  </BBTable>
</template>

<script lang="ts">
import { reactive, PropType } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import { BBTableColumn } from "../bbkit/types";
import { urlfy, instanceSlug, environmentName } from "../utils";
import { EnvironmentId, Instance } from "../types";

interface LocalState {
  columnList: BBTableColumn[];
  dataSource: Object[];
}

export default {
  name: "InstanceTable",
  components: {},
  props: {
    instanceList: {
      required: true,
      type: Object as PropType<Instance[]>,
    },
  },
  setup(props, ctx) {
    const store = useStore();

    const state = reactive<LocalState>({
      columnList: [
        {
          title: "",
        },
        {
          title: "Name",
        },
        {
          title: "Environment",
        },
        {
          title: "Host:Port",
        },
        {
          title: "External link",
        },
        {
          title: "Created",
        },
      ],
      dataSource: [],
    });

    const router = useRouter();

    const clickInstance = function (section: number, row: number) {
      const instance = props.instanceList[row];
      router.push(`/instance/${instanceSlug(instance)}`);
    };

    const environmentNameFromId = function (id: EnvironmentId) {
      return environmentName(store.getters["environment/environmentById"](id));
    };

    return {
      state,
      urlfy,
      clickInstance,
      environmentNameFromId,
    };
  },
};
</script>
