const usageCtx = document.getElementById("usageChart").getContext("2d");
const usageChart = new Chart(usageCtx, {
  type: "bar",
  data: {
    labels: ["Jan", "Feb", "Mar", "Apr", "Mei", "Jun"],
    datasets: [
      {
        label: "Gedung A",
        data: [25, 30, 22, 28, 32, 27],
        backgroundColor: "rgba(59, 130, 246, 0.6)",
      },
      {
        label: "Gedung B",
        data: [18, 25, 20, 22, 27, 24],
        backgroundColor: "rgba(16, 185, 129, 0.6)",
      },
      {
        label: "Gedung C",
        data: [12, 15, 17, 20, 18, 22],
        backgroundColor: "rgba(245, 158, 11, 0.6)",
      },
    ],
  },
  options: {
    responsive: true,
    maintainAspectRatio: false,
    scales: {
      y: {
        beginAtZero: true,
        title: {
          display: true,
          text: "Jumlah Reservasi",
        },
      },
    },
  },
});

// Status Chart
const statusCtx = document.getElementById("statusChart").getContext("2d");
const statusChart = new Chart(statusCtx, {
  type: "pie",
  data: {
    labels: ["Disetujui", "Ditolak", "Menunggu"],
    datasets: [
      {
        data: [70, 15, 15],
        backgroundColor: [
          "rgba(16, 185, 129, 0.7)",
          "rgba(239, 68, 68, 0.7)",
          "rgba(245, 158, 11, 0.7)",
        ],
        borderColor: [
          "rgba(16, 185, 129, 1)",
          "rgba(239, 68, 68, 1)",
          "rgba(245, 158, 11, 1)",
        ],
        borderWidth: 1,
      },
    ],
  },
  options: {
    responsive: true,
    maintainAspectRatio: false,
  },
});
