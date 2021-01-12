import MajorChart from './Major';
import './Chart.css';

const [h1,h2,h3] = ["공부 시간","가장 사랑하는 프로그래밍 언어","공부 방법"];

const Chart = (props) => {
    return(
        <div className="chart">
            <MajorChart title={ h1 }/>
            <MajorChart title={ h2 }/>
            <MajorChart title={ h3 }/>
        </div>
    );
}

export default Chart;