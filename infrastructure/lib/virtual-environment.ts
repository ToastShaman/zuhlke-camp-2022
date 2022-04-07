export class VirtualEnvironment {
    constructor(readonly name: string) { }
    
    toString() {
        return this.name;
    }

    withPrefix(prefix: string): VirtualEnvironment {
        return new VirtualEnvironment(`${prefix}-${this.name}`.toLowerCase());
    }

    static readonly CI = new VirtualEnvironment("ci")
    static readonly QA = new VirtualEnvironment("qa")
    static readonly FNCTNL = new VirtualEnvironment("fnctnl")
    static readonly STAGING = new VirtualEnvironment("staging")

    static readonly DEV = [
        VirtualEnvironment.CI,
        VirtualEnvironment.QA,
        VirtualEnvironment.FNCTNL
    ]

    static readonly PRE_PROD = [
        VirtualEnvironment.STAGING
    ]
};