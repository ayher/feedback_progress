import './up_down.css'
function UpDown(props){
    return <div className="up_down">
        <div className="up">
            {props.up}
        </div>
        <div className="down">
            {props.down}
        </div>
    </div>
}

export default UpDown;
