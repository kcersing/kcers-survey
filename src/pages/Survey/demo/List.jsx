import React from 'react';
import { ProTable, ProCard } from '@ant-design/pro-components';
import { FileAddOutlined } from '@ant-design/icons';
import { listSurvey, deleteSurvey } from '@/services/ant-design-pro/survey';
import { history } from 'umi';
import { message } from 'antd';
const SurveyList = () => {
    const columns = [
        {
            title: 'ID',
            dataIndex: 'id',
            valueType: 'digit',
        },
        {
            title: '标题',
            dataIndex: 'title',
            render: (text, record) => (<a onClick={() => history.push(`/survey/${record.id}/design`)}>{text}</a>),
        },
        {
            title: '状态',
            dataIndex: 'status',
            valueEnum: {
                draft: { text: '草稿', status: 1 },
                active: { text: '发布中', status: 2 },
                closed: { text: '已关闭', status: 0 },
            },
        },
        {
            title: '创建时间',
            dataIndex: 'createdAt',
            valueType: 'dateTime',
        },
        {
            title: '操作',
            valueType: 'option',
            render: (_, record) => [
                <a onClick={() => history.push(`/survey/${record.id}/edit`)}>编辑</a>,
                <a onClick={() => history.push(`/survey/${record.id}/design`)}>设计</a>,
                <a onClick={() => history.push(`/survey/${record.id}/respond`)}>预览</a>,
                <a onClick={() => history.push(`/survey/${record.id}/statistics`)}>统计</a>,
                <a danger onClick={() => {
                        deleteSurvey(record.id).then(() => {
                            message.success('删除成功');
                            tableRef.current?.reload();
                        });
                    }}>
          删除
        </a>,
            ],
        },
    ];
    const tableRef = React.createRef();
    return (<ProCard title="问卷列表" bordered={false} headerExtra={<a onClick={() => history.push('/survey/create')}>
          <FileAddOutlined /> 创建问卷
        </a>}>
      <ProTable ref={tableRef} rowKey="id" request={listSurvey} columns={columns} toolBarRender={() => []} pagination={{
            pageSize: 10,
        }}/>
    </ProCard>);
};
export default SurveyList;
//# sourceMappingURL=List.jsx.map