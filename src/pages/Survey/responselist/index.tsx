import { EllipsisOutlined, PlusOutlined } from '@ant-design/icons';
import type { ActionType, ProColumns } from '@ant-design/pro-components';
import { ProTable, TableDropdown } from '@ant-design/pro-components';
import { Button, Dropdown, Input, Space, Tag } from 'antd';
import React, { useRef } from 'react';
import request from 'umi-request';
import { history } from '@@/core/history';
import { listSurvey ,listResponse} from '@/services/ant-design-pro/survey';
import { useParams } from '@@/exports';
export const waitTimePromise = async (time: number = 100) => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve(true);
    }, time);
  });
};

export const waitTime = async (time: number = 100) => {
  await waitTimePromise(time);
};

const columns: ProColumns<API.Response>[] = [
  {
    title: '编号',
    dataIndex: 'sn',
    valueType: 'textarea',
    ellipsis: true,
  },
  {
    title:'受访人',
    dataIndex: 'respondent',
    ellipsis: true,
    tip: '受访人',
  },
  {
    title:'受访人联系电话',
    dataIndex: 'respondentPhone',
    ellipsis: true,
    tip: '受访人',
  },
  {
    title:'调研员',
    dataIndex: 'researcher',
    ellipsis: true,
    tip: '受访人',
  },
  {
    title:'调研员',
    dataIndex: 'researcherPhone',
    ellipsis: true,
    tip: '受访人',
  },
  {
    title: '填写问卷时间',
    sorter: true,
    dataIndex: 'createdAt',

  },
  {
    title:'完成度',
    dataIndex: 'answerCount',
    ellipsis: true,
    tip: '如100题，填写了80道题，完成度为80',
  },
  {
    title: '操作',
    dataIndex: 'option',
    valueType: 'option',
    render: (_, record) => [
      <a key="config" onClick={() => { history.push(`#`)} }>详情</a>,
      // <a key="config" onClick={() => { history.push(`/survey/${record.id}/statistics`)}}>统计</a>,
      <a key="remove" onClick={() => { history.push(`#`)} }>设置状态</a>,

    ],
  },
];

export default () => {


  const { id } = useParams();
  const surveyId = id ? parseInt(id) : 1;
console.log(surveyId)

  const actionRef = useRef<ActionType>();
  return (
    <ProTable<API.Response>
      columns={columns}
      actionRef={actionRef}
      cardBordered
      // request={listResponse}
      request={async (params, sort, filter) => {

        params.surveyId = surveyId;
        console.log(params)
        console.log(sort, filter);
        return  listResponse({...params})
        // await waitTime(2000);

      }}
      // editable={{
      //   type: 'multiple',
      // }}
      // columnsState={{
      //   persistenceKey: 'pro-table-singe-demos',
      //   persistenceType: 'localStorage',
      //   defaultValue: {
      //     option: { fixed: 'right', disable: true },
      //   },
      //   onChange(value) {
      //     console.log('value: ', value);
      //   },
      // }}
      rowKey="id"
      search={{
        labelWidth: 'auto',
      }}
      options={{
        setting: {
          listsHeight: 400,
        },
      }}
      // form={{
      //   // Since transform is configured, the submitted parameters are different from the defined ones, so they need to be transformed here
      //   syncToUrl: (values, type) => {
      //     if (type === 'get') {
      //       return {
      //         ...values,
      //         created_at: [values.startTime, values.endTime],
      //       };
      //     }
      //     return values;
      //   },
      // }}
      // pagination={{
      //   pageSize: 10,
      //   onChange: (page) => console.log(page),
      // }}
      dateFormatter="string"
      headerTitle="已填写问卷列表"
      toolBarRender={() => [
        <Button
          key="button"
          onClick={() => {
            actionRef.current?.reload();
          }}
          type="primary"
        >
          导出
        </Button>,
      ]}
    />
  );
};
