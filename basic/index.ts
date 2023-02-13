import * as pulumi from "@pulumi/pulumi";
import * as k8s from "@pulumi/kubernetes";
import * as kx from "@pulumi/kubernetesx";

const config = new pulumi.Config();
const tag = config.require("tag");

const appLabels = { app: "nginx" };
const deployment = new k8s.apps.v1.Deployment("nginx", {
    spec: {
        selector: { matchLabels: appLabels },
        replicas: 3,
        template: {
            metadata: { labels: appLabels },
            spec: { containers: [{ name: "nginx", image: pulumi.interpolate `nginx:${tag}` }] }
        }
    }
});
export const name = deployment.metadata.name;
