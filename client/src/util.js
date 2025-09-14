import moment from "moment";

export default function formatData(data) {
  const formatedData = moment(data).startOf("second").fromNow();
  return formatedData;
}
