import React, {FC} from 'react';
import './App.less';
import {Layout} from 'antd';
import {Content, Header} from "antd/es/layout/layout";
import {ConvertImage} from "./ConvertImage";

const App: FC = () => (
    <Layout>
        <Header>
            <div className="logo"/>
            <h1 className="page-title">在线图像文件转换器</h1>
        </Header>
        <Content className="content">
            <ConvertImage/>
        </Content>
    </Layout>
);

export default App;
