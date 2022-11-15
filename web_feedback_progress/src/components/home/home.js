
import './home.css'
import RankItem from "./rank_item";

function Home() {
    const RankData=[
        {"rank":"一","name":"默默","title":"给那个焦虑的我","interact":{"love":100,"comment":23,"collect":26}},
        {"rank":"二","name":"默默","title":"给那个焦虑的我","interact":{"love":100,"comment":23,"collect":26}},
        {"rank":"三","name":"默默","title":"给那个焦虑的我","interact":{"love":100,"comment":23,"collect":26}},
    ]

    var rankItems = []
    for (let i=0;i<RankData.length;i++){
        rankItems.push(<RankItem{...RankData[i]} />)
    }

    return (
        <div className="home">
            {rankItems}
        </div>
    );
}

export default Home;
