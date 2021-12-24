export function toast(that, msg, type , pos , delay) {


  // unicode symbols
  const infoSymbol = "\u2139"
  const successSymbol = "\u2713"
  const warningSymbol = "\u26A0"
  const errorSymbol = "\u2715"


  let msgClass = []
  let actionClass = []
  let symbol = ''

  if (type === "info") {

    msgClass.push("text-gray-200", "py-4", "bg-gray-800")
    actionClass.push("text-blue-600", "bg-gray-900")
    symbol = infoSymbol

  } else if (type === "success") {

    msgClass.push("text-gray-200", "py-4", "bg-green-700")
    actionClass.push("text-blue-800", "bg-draculaGreen")
    symbol = successSymbol

  } else if (type === "warning") {

    msgClass.push("text-gray-50", "py-4", "bg-yellow-400")
    actionClass.push("text-yellow-600", "bg-yellow-800")
    symbol = warningSymbol

  } else {

    msgClass.push("text-gray-200", "py-4", "bg-red-600")
    actionClass.push("text-red-700", "bg-red-900")
    symbol = errorSymbol
  }
  that.$toasted.show(msg, {
    theme: "none",
    position: pos,
    className: msgClass,
    duration: delay,
    action: {
      text: symbol,
      onClick: (e, toastObject) => {
        toastObject.goAway(0);
      },
      class: actionClass
    },
  });
}

