import { MailOutlined } from '@ant-design/icons';
import type { ProColumns } from '@ant-design/pro-components';
import { ProTable } from '@ant-design/pro-components';
import { Card, Descriptions, Menu } from 'antd';
import { useState,useEffect  } from 'react';
import {getSurvey, listQuestion, treeQuestion} from '@/services/ant-design-pro/survey';
import {useParams} from "react-router";
import { DemoCustomColor } from '@/pages/survey/statistics/components/custom-color';
import { DemoRose } from '@/pages/survey/statistics/components/donut-rose';
import { DemoMemo } from '@/pages/survey/statistics/components/memo';
import { Demobase } from '@/pages/survey/statistics/components/space-layer';
import { DemoPie } from '@/pages/survey/statistics/components/spider-label';
import { DemoDendrogram } from '@/pages/survey/statistics/components/vertical-tidy-tree';

const waitTime = (time: number = 100) => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve(true);
    }, time);
  });
};

export type TableListItem = {
  key: number;
  name: string;
  createdAt: number;
  progress: number;
};
const tableListDataSource: TableListItem[] = [];

for (let i = 0; i < 2; i += 1) {
  tableListDataSource.push({
    key: i,
    name: `TradeCode ${i}`,
    createdAt: Date.now() - Math.floor(Math.random() * 2000),
    progress: Math.ceil(Math.random() * 100) + 1,
  });
}

const columns: ProColumns<TableListItem>[] = [
  {
    title: '序号',
    dataIndex: 'index',
    valueType: 'index',
    width: 80,
  },
  {
    title: '更新时间',
    key: 'since2',
    dataIndex: 'createdAt',
    valueType: 'date',
  },
  {
    title: '执行进度',
    dataIndex: 'progress',
    valueType: 'progress',
  },
];

export default () => {
  const [key, setKey] = useState('1');
  const [survey, setSurvey] = useState<API.Survey>({});
  const [menuItems, setMenuItems] = useState<Menu.ItemType[]>([]);

  let params = useParams();
  const surveyId =parseInt(params.id)
  useEffect(() => {
    const getSurveyInfo = async () => {
      const surveyData = await getSurvey({id: surveyId})
      setSurvey(surveyData.data);
    }
    getSurveyInfo()
  }, []);

  console.log(survey)

  useEffect(() => {
    const fetchMenuItems = async () => {
      try {
        const response = await treeQuestion({ surveyId: 2 }); // 根据实际情况传入参数
        const items = convertToMenuItems(response.data);
        console.log(response.data);
        setMenuItems(items);
      } catch (error) {
        console.error('获取菜单数据失败:', error);
      }
    };

    fetchMenuItems();
  }, []);

  const convertToMenuItems = (data: any[]): Menu.ItemType[] => {
    return data.map(item => ({
      key: item.value,
      label: item.title,
      title: item.title,
      // 如果有子节点，递归转换
      children: item.children ? convertToMenuItems(item.children) : undefined,
    }));
  };


  return (
    <Card>


    <ProTable<TableListItem>
      columns={columns}
      rowKey="key"
      pagination={{
        showSizeChanger: true,
      }}
      tableRender={(_, dom) => (
        <div
          style={{
            display: 'flex',
            width: '100%',
          }}
        >
          <Menu
            onSelect={(e) => setKey(e.key as string)}
            style={{ width: 256 }}
            defaultSelectedKeys={['1']}
            defaultOpenKeys={['sub1']}
            mode="inline"
            items={menuItems}
          />
          <div
            style={{
              flex: 1,
            }}
          >
            {dom}
          </div>
        </div>
      )}
      tableExtraRender={(_, data) => (
        <Card>
         <h3> {survey.title?survey.title:""}</h3>
          <Descriptions size="small" column={3}>
            <Descriptions.Item label="已完成问卷数量">69682</Descriptions.Item>
            <Descriptions.Item label="共完成题目">586246</Descriptions.Item>
            <Descriptions.Item label="受访人">
              69682
            </Descriptions.Item>
            <Descriptions.Item label="调研员">
              500
            </Descriptions.Item>
            <Descriptions.Item label="走访村庄">
              2658
            </Descriptions.Item>
          </Descriptions>


          <DemoCustomColor />
          <DemoRose />
          <DemoMemo />
          <Demobase />
          <DemoPie />
          <DemoDendrogram />




        </Card>
      )}
      params={{
        key,
      }}
      request={async () => {
        await waitTime(200);
        return {
          success: true,
          data: tableListDataSource,
        };
      }}
      dateFormatter="string"
      headerTitle="自定义表格主体"
    />
    </Card>
  );
};
