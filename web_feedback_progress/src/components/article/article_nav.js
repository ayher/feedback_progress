import { DoubleRightOutlined } from '@ant-design/icons';
import './article.css'

function ArticleNav(props){
    return <>
        <div className="article_nav_inner">{props[0]}</div>
        <DoubleRightOutlined className="article_nav_icon"/>
        <div className="article_nav_inner">{props[1]}</div>
        <DoubleRightOutlined className="article_nav_icon"/>
        <div className="article_nav_inner">{props[2]}</div>
    </>
}

export default ArticleNav;