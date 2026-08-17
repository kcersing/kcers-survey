import { interviewerLogSearch } from '@/services/ant-design-pro/survey';
import { LogoutOutlined, SearchOutlined, UserOutlined } from '@ant-design/icons';
import { PageContainer } from '@ant-design/pro-components';
import { history } from '@umijs/max';
import {
  Button,
  Card,
  Descriptions,
  Empty,
  Input,
  message,
  Modal,
  Space,
  Table,
  Tag,
  Typography,
} from 'antd';
import type { ColumnsType } from 'antd/es/table';
import React, { useEffect, useState } from 'react';

const { Title } = Typography;

interface LogEntry {
  id: number;
  api: string;
  method: string;
  success: boolean;
  reqContent: string;
  respContent: string;
  ip: string;
  userAgent: string;
  operatorsr: string;
  time: number;
  createdAt: string;
}

const columns: ColumnsType<LogEntry> = [
  { title: 'ID', dataIndex: 'id', key: 'id', width: 70 },
  { title: '接口', dataIndex: 'api', key: 'api', width: 260, ellipsis: true },
  { title: '方法', dataIndex: 'method', key: 'method', width: 80 },
  {
    title: '状态',
    dataIndex: 'success',
    key: 'success',
    width: 80,
    render: (v: boolean) => (v ? <Tag color="green">成功</Tag> : <Tag color="red">失败</Tag>),
  },
  { title: 'IP', dataIndex: 'ip', key: 'ip', width: 140 },
  { title: '操作者', dataIndex: 'operatorsr', key: 'operatorsr', width: 100 },
  { title: '耗时(ms)', dataIndex: 'time', key: 'time', width: 90 },
  { title: '时间', dataIndex: 'createdAt', key: 'createdAt', width: 170 },
  {
    title: '操作',
    key: 'action',
    width: 80,
    fixed: 'right' as const,
    render: (_, record) => (
      <Button
        type="link"
        size="small"
        onClick={(e) => {
          e.stopPropagation();
          Modal.info({
            title: '请求详情',
            width: 700,
            content: (
              <Descriptions column={1} size="small" style={{ marginTop: 16 }}>
                <Descriptions.Item label="接口">{record.api}</Descriptions.Item>
                <Descriptions.Item label="方法">{record.method}</Descriptions.Item>
                <Descriptions.Item label="IP">{record.ip}</Descriptions.Item>
                <Descriptions.Item label="UA">{record.userAgent}</Descriptions.Item>
                <Descriptions.Item label="请求内容">
                  <pre style={{ maxHeight: 300, overflow: 'auto', fontSize: 12 }}>
                    {formatJSON(record.reqContent)}
                  </pre>
                </Descriptions.Item>
                <Descriptions.Item label="响应内容">
                  <pre style={{ maxHeight: 300, overflow: 'auto', fontSize: 12 }}>
                    {formatJSON(record.respContent)}
                  </pre>
                </Descriptions.Item>
              </Descriptions>
            ),
          });
        }}
      >
        详情
      </Button>
    ),
  },
];

function formatJSON(s: string): string {
  if (!s) return '-';
  try {
    return JSON.stringify(JSON.parse(s), null, 2);
  } catch {
    return s;
  }
}

const InterviewerLogs: React.FC = () => {
  const [loading, setLoading] = useState(false);
  const [data, setData] = useState<LogEntry[]>([]);
  const [total, setTotal] = useState(0);
  const [pagination, setPagination] = useState({ current: 1, pageSize: 20 });
  const [keyword, setKeyword] = useState('');
  const [searched, setSearched] = useState(false);
  const [mobile, setMobile] = useState('');

  const token = sessionStorage.getItem('interviewer_token');
  const savedMobile = sessionStorage.getItem('interviewer_mobile');

  useEffect(() => {
    if (!token) {
      history.push('/interviewer/login');
      return;
    }
    setMobile(savedMobile || '');
  }, []);

  const handleSearch = async (page = 1, pageSize = pagination.pageSize) => {
    if (!token || !keyword.trim()) return;
    setLoading(true);
    try {
      const res = await interviewerLogSearch(
        { keyword: keyword.trim(), current: page, pageSize },
        token,
      );
      if (res.code === 0) {
        setData(res.data || []);
        setTotal(res.total || 0);
        setPagination({ current: page, pageSize });
        setSearched(true);
      } else if (res.code === 10002) {
        message.error('登录已过期，请重新登录');
        handleLogout();
      } else {
        message.error(res.msg || '查询失败');
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
    handleSearch(pag.current || 1, pag.pageSize || 20);
  };

  return (
    <PageContainer header={{ title: '' }}>
      <div style={{ maxWidth: 1200, margin: '0 auto', padding: '40px 0' }}>
        <div
          style={{
            display: 'flex',
            justifyContent: 'space-between',
            alignItems: 'center',
            marginBottom: 24,
          }}
        >
          <Title level={3} style={{ margin: 0 }}>
            提交日志查询
          </Title>
          <Space>
            <span>
              <UserOutlined /> {mobile}
            </span>
            <Button
              icon={<LogoutOutlined />}
              onClick={handleLogout}
              type="link"
            >
              退出登录
            </Button>
          </Space>
        </div>

        <Card>
          <Space style={{ marginBottom: 16 }}>
            <Input
              placeholder="输入手机号、IP 或地址进行搜索"
              value={keyword}
              onChange={(e) => setKeyword(e.target.value)}
              onPressEnter={() => handleSearch(1)}
              style={{ width: 320 }}
              prefix={<SearchOutlined />}
              allowClear
            />
            <Button type="primary" onClick={() => handleSearch(1)} loading={loading}>
              搜索
            </Button>
          </Space>

          <Table<LogEntry>
            columns={columns}
            dataSource={data}
            rowKey="id"
            loading={loading}
            scroll={{ x: 1100 }}
            pagination={{
              current: pagination.current,
              pageSize: pagination.pageSize,
              total,
              showSizeChanger: true,
              showTotal: (t) => `共 ${t} 条记录`,
            }}
            onChange={handleTableChange}
            locale={{
              emptyText: searched ? <Empty description="未找到相关日志" /> : <Empty description="请输入关键字搜索" />,
            }}
          />
        </Card>
      </div>
    </PageContainer>
  );
};

export default InterviewerLogs;