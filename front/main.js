// function runJob(jobId) {
//     const params = new URLSearchParams({
//         action: "run",
//         job_id: jobId
//     });
//
//     fetch("/cron-manager?" + params.toString(), {
//         method: "GET",
//     })
//         .then(response => {
//             if (response.ok) {
//                 alert("Job started successfully!");
//             } else {
//                 alert("Failed to start job");
//             }
//         })
//         .catch(err => {
//             alert("Error: " + err);
//         });
// }