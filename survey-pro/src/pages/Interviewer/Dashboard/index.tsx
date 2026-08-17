import { interviewerSurveyList } from '@/services/ant-design-pro/survey';
import { KeyOutlined, LogoutOutlined, SearchOutlined, UserOutlined } from '@ant-design/icons';
import { PageContainer } from '@ant-design/pro-components';
import { history } from '@umijs/max';
import { Button, Card, Empty, message, Space, Table, Typography } from 'antd';
import type { ColumnsType } from 'antd/es/table';
import React, { useEffect, useState } from 'react';

const { Title } = Typography;

interface SurveyItem {
  survey_id: number;
  survey_title: string;
  sn: string;
  respondent: string;
  respondent_phone: string;
  address: string;
  created_at: string;
  answers_count: number;
}

const columns: ColumnsType<SurveyItem> = [
  { title: '问卷标题', dataIndex: 'survey_title', key: 'survey_title', width: 200 },
  { title: '受访者', dataIndex: 'respondent', key: 'respondent', width: 100 },
  { title: '受访者电话', dataIndex: 'respondent_phone', key: 'respondent_phone', width: 130 },
  { title: '调研地址', dataIndex: 'address', key: 'address', width: 200 },
  { title: '提交时间', dataIndex: 'created_at', key: 'created_at', width: 170 },
  { title: '答题数', dataIndex: 'answers_count', key: 'answers_count', width: 80 },
];

const InterviewerDashboard: React.FC = () => {
  const [loading, setLoading] = useState(false);
  const [data, setData] = useState<SurveyItem[]>([]);
  const [total, setTotal] = useState(0);
  const [pagination, setPagination] = useState({ current: 1, pageSize: 10 });
  const [mobile, setMobile] = useState('');

  const token = sessionStorage.getItem('interviewer_token');
  const savedMobile = sessionStorage.getItem('interviewer_mobile');

  useEffect(() => {
    if (!token) {
      history.push('/interviewer/login');
      return;
    }
    setMobile(savedMobile || '');
    fetchSurveys(1, 10);
  }, []);

  const fetchSurveys = async (page: number, pageSize: number) => {
    if (!token) return;
    setLoading(true);
    try {
      const res = await interviewerSurveyList({ current: page, pageSize }, token);
      if (res.code === 0) {
        setData(res.data || []);
        setTotal(res.total || 0);
        setPagination({ current: page, pageSize });
      } else if (res.code === 10002) {
        message.error('登录已过期，请重新登录');
        handleLogout();
      }
    } catch {
      message.error('查询失败');
    } finally {
      setLoading(false);
    }
  };

  const handleLogout = () => {
    sessionStorage.removeItem('interviewer_token');
    sessionStorage.removeItem('interviewer_mobile');
    history.push('/interviewer/login');
  };

  const handleTableChange = (pag: { current?: number; pageSize?: number }) => {
    fetchSurveys(pag.current || 1, pag.pageSize || 10);
  };

  return (
    <PageContainer header={{ title: '' }}>
      <div style={{ maxWidth: 1000, margin: '0 auto', padding: '40px 0' }}>
        <div
          style={{
            display: 'flex',
            justifyContent: 'space-between',
            alignItems: 'center',
            marginBottom: 24,
          }}
        >
          <Title level={3} style={{ margin: 0 }}>
            我的调研问卷
          </Title>
          <Space>
            <span>
              <UserOutlined /> {mobile}
            </span>
            <Button
              icon={<SearchOutlined />}
              type="link"
              onClick={() => history.push('/interviewer/logs')}
            >
              日志查询
            </Button>
            <Button
              icon={<KeyOutlined />}
              type="link"
              onClick={() => history.push('/interviewer/change-password')}
            >
              修改密码
            </Button>
            <Button icon={<LogoutOutlined />} onClick={handleLogout} type="link">
              退出登录
            </Button>
          </Space>
        </div>

        <Card>
          <Table<SurveyItem>
            columns={columns}
            dataSource={data}
            rowKey="sn"
            loading={loading}
            scroll={{ x: 900 }}
            onRow={(record) => ({
              onClick: () => {
                window.open(`/survey/${record.survey_id}/response/${record.sn}?interviewer=1`);
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
              emptyText: <Empty description="暂无调研问卷" />,
            }}
          />
        </Card>
      </div>
    </PageContainer>
  );
};

export default InterviewerDashboard;
