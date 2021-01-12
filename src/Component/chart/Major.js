import { Chart } from "react-google-charts";
import './Major.css';

const MajorChart = ( {title} ) => {
  return (
    <div>
      <Chart
      width={'380px'}
      height={'540px'}
      chartType="PieChart"
      loader={<div>. </div>}
      data={[
        ['Task', 'Hours per Day'],
        ['Work', 11],
        ['Eat', 2],
        ['Commute', 2],
        ['Watch TV', 2],
        ['Sleep', 7],
      ]}
      options={{
        title: { title },
      }}
      rootProps={{ 'data-testid': '1' }}
    />
    <hr className="hr"/>
    </div>
  );
};

export default MajorChart;