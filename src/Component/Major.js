import { Chart } from "react-google-charts";

const MajorChart = () => {
  return (
    <Chart
      width={'700px'}
      height={'500px'}
      chartType="PieChart"
      loader={<div></div>}
      data={[
        ['Task', 'Hours per Day'],
        ['Work', 11],
        ['Eat', 2],
        ['Commute', 2],
        ['Watch TV', 2],
        ['Sleep', 7],
      ]}
      options={{
        title: '전공 분야',
      }}
      rootProps={{ 'data-testid': '1' }}
    />
  );
};

export default MajorChart;