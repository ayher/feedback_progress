import React, { useEffect } from 'react'
import UpDown from "../../public/updown_layout/up_down";
import DownNav from "../../public/down_nav/down_nav";
import './home.css'
import RankItem from "./rank_item";
import {getArticleInfo, getRankArticleInfo} from "../../redux/action/article";
import { useSelector,useDispatch } from "react-redux";


function Home() {
    var dispatch = useDispatch()
    useEffect(() => {
        dispatch(getRankArticleInfo())      //写在函数组件里
      }, [])    

    const article = useSelector(state => state.Article);
    const RankData=article.ArticleRand ?? []

    var rankItems = []
    for (let i=0;i<RankData.length;i++){
        RankData[i].rank = i + 1
        rankItems.push(<RankItem{...RankData[i]} key={i}/>)
    }

    var nav = [<div>登陆</div>,<div>去创作</div>,<div>其他排名</div>];

    return (
        <UpDown{...{
            "up":
                <div className="home">
                    {rankItems}
                </div>
            ,"down":
                <DownNav{...{nav}} />
        }}/>
    );
}

export default Home;
