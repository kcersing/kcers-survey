import { queryByPhone } from '@/services/ant-design-pro/survey';
import { InfoCircleOutlined, SearchOutlined } from '@ant-design/icons';
import { PageContainer } from '@ant-design/pro-components';
import { Card, Empty, Input, message, Table, Typography } from 'antd';
import type { ColumnsType } from 'antd/es/table';
import React, { useState } from 'react';

const { Title } = Typography;

interface QueryResult {
  survey_id: number;
  survey_title: string;
  sn: string;
  respondent: string;
  respondent_phone: string;
  address: string;
  created_at: string;
  answers_count: number;
}

const columns: ColumnsType<QueryResult> = [
  { title: '问卷标题', dataIndex: 'survey_title', key: 'survey_title', width: 200 },
  { title: '受访者', dataIndex: 'respondent', key: 'respondent', width: 100 },
  { title: '受访者电话', dataIndex: 'respondent_phone', key: 'respondent_phone', width: 130 },
  { title: '调研地址', dataIndex: 'address', key: 'address', width: 200 },
  { title: '提交时间', dataIndex: 'created_at', key: 'created_at', width: 170 },
  { title: '答题数', dataIndex: 'answers_count', key: 'answers_count', width: 80 },
];

const QueryPage: React.FC = () => {
  const [loading, setLoading] = useState(false);
  const [data, setData] = useState<QueryResult[]>([]);
  const [total, setTotal] = useState(0);
  const [phone, setPhone] = useState('');
  const [searched, setSearched] = useState(false);
  const [pagination, setPagination] = useState({ current: 1, pageSize: 10 });

  const handleSearch = (value: string, page = 1, pageSize = 10) => {
    const trimmed = value.trim();
    if (!trimmed) {
      message.warning('请输入手机号');
      return;
    }
    setPhone(trimmed);
    setLoading(true);
    queryByPhone({ phone: trimmed, current: page, pageSize })
      .then((res) => {
        setData(res?.data || []);
        setTotal(res?.total || 0);
        setSearched(true);
        setPagination({ current: page, pageSize });
      })
      .catch(() => {
        message.error('查询失败，请稍后重试');
      })
      .finally(() => {
        setLoading(false);
      });
  };

  const handleTableChange = (pag: { current?: number; pageSize?: number }) => {
    if (phone) {
      handleSearch(phone, pag.current || 1, pag.pageSize || 10);
    }
  };

  return (
    <PageContainer header={{ title: '' }}>
      <div style={{ maxWidth: 1000, margin: '0 auto', padding: '40px 0' }}>
        <Title level={3} style={{ textAlign: 'center', marginBottom: 32 }}>
          调研员问卷查询
        </Title>

        <Card style={{ marginBottom: 24 }}>
          <Input.Search
            size="large"
            placeholder="请输入调研员手机号"
            enterButton={<><SearchOutlined /> 查询</>}
            onSearch={(value) => handleSearch(value)}
            loading={loading}
            maxLength={11}
            allowClear
          />
        </Card>

        {!searched ? (
          <Card>
            <Empty description="请输入手机号查询" image={Empty.PRESENTED_IMAGE_SIMPLE} />
          </Card>
        ) : (
          <Card>
            <Table<QueryResult>
              columns={columns}
              dataSource={data}
              rowKey="sn"
              loading={loading}
              scroll={{ x: 900 }}
              onRow={(record) => ({
                onClick: () => {
                  window.open(`/survey/${record.survey_id}/response/${record.sn}`);
                },
                style: { cursor: 'pointer' },
              })}
              pagination={{
                current: pagination.current,
                pageSize: pagination.pageSize,
                total,
                showSizeChanger: true,
                showTotal: (t) => `共 ${t} 条记录`,
              }}
              onChange={handleTableChange}
              locale={{
                emptyText: <Empty description="未查询到相关问卷" />,
              }}
            />
          </Card>
        )}
      </div>
    </PageContainer>
  );
};

export default QueryPage;