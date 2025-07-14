
import type { ProColumns } from '@ant-design/pro-components';
import { ProTable } from '@ant-design/pro-components';
import { Card, Descriptions, Menu, Modal,Button } from 'antd';
import { useState,useEffect  } from 'react';
import {getSurvey, treeQuestion,getResponseAnswers,heatmap,questionBasicData} from '@/services/ant-design-pro/survey';
import {useParams} from "react-router";
import { DemoCustomColor } from '@/pages/survey/statistics/components/custom-color';
import { DemoRose } from '@/pages/survey/statistics/components/donut-rose';
import { Demobase } from '@/pages/survey/statistics/components/space-layer';
import { DemoPie } from '@/pages/survey/statistics/components/spider-label';
import {HeatMap} from "@/pages/survey/statistics/components/heatmap";

const columns: ProColumns[] = [
  {
    title: '问题',
    dataIndex: 'content',
    width: 80,
  },
  {
    title: '回答',
    dataIndex: 'answer',
    width: 80,
  },
  {
    title: '补充',
    dataIndex: 'answerText',
    width: 80,
  },

  // 回答内容:{item.answer.map(v=>{ return (<a>{v}；</a>)} )}


  {
    title: '创建时间',
    key: 'since2',
    dataIndex: 'createdAt',
    valueType: 'date',
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


  const [openheatmap, setOpenheatmap] = useState<boolean>(false);
  const showLoading = () => {


    setOpenheatmap(true);

  };



  return (
    <>
    <Card>
    <ProTable<API.ResponseAnswers>
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

          <a type="primary" onClick={showLoading}>
            地图
          </a>

        </Card>
      )}
      params={{
        key,
      }}
      request={async () => {

        console.log(key)

        const  ans = await getResponseAnswers({id:parseInt(key)} );
        return {
          success: true,
          data: ans.data,
        };
      }}
      dateFormatter="string"
      headerTitle="自定义表格主体"
    />

    </Card>

  <Modal
    title={<p>地图</p>}
    centered
    width={{
      xs: '90%',
      sm: '80%',
      md: '70%',
      lg: '70%',
      xl: '70%',
      xxl: '80%',
    }}
    footer={<></>}
    open={openheatmap}
    onCancel={() => setOpenheatmap(false)}
  >
    <HeatMap data={"1"}/>
  </Modal>

  {/*<DemoCustomColor data={} />*/}
  {/*<DemoRose data={}/>*/}
  {/*<Demobase data={}/>*/}
  {/*<DemoPie data={}/>*/}

 </>



  );
};
